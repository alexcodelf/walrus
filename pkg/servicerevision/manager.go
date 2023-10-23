package servicerevision

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/terraform/parser"

	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
)

const (
	// DeployerType of the revision.
	DeployerType = types.DeployerTypeTF

	JobTypeApply   = "apply"
	JobTypeDestroy = "destroy"

	// _backendAPI the API path to terraform deploy backend.
	// Terraform will get and update deployment states from this API.
	_backendAPI = "/v1/projects/%s/environments/%s/services/%s/revisions/%s/terraform-states"
)

type Manager struct {
	modelClient model.ClientSet
}

func NewRevisionManager(mc model.ClientSet) *Manager {
	return &Manager{
		modelClient: mc,
	}
}

type CreateOptions struct {
	// ServiceID indicates the ID of service which is for create the revision.
	ServiceID object.ID
	// WorkflowStepExecutionID indicates the ID of workflow step execution which is for create the revision.
	WorkflowStepExecutionID object.ID
	// JobType indicates the type of the job.
	JobType string
}

func (m Manager) Create(ctx context.Context, opts CreateOptions) (*model.ServiceRevision, error) {
	// Validate if there is a running revision.
	prevEntity, err := m.modelClient.ServiceRevisions().Query().
		Where(servicerevision.And(
			servicerevision.ServiceID(opts.ServiceID),
			servicerevision.DeployerType(DeployerType))).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if prevEntity != nil && status.ServiceRevisionStatusRunning.IsUnknown(prevEntity) {
		return nil, errors.New("service deployment is running")
	}

	// Get the corresponding service and template version.
	svc, err := m.modelClient.Services().Query().
		Where(service.ID(opts.ServiceID)).
		WithTemplate(func(tvq *model.TemplateVersionQuery) {
			tvq.Select(
				templateversion.FieldName,
				templateversion.FieldVersion)
		}).
		Select(
			service.FieldID,
			service.FieldProjectID,
			service.FieldEnvironmentID,
			service.FieldAttributes).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	entity := &model.ServiceRevision{
		ProjectID:               svc.ProjectID,
		EnvironmentID:           svc.EnvironmentID,
		ServiceID:               svc.ID,
		TemplateName:            svc.Edges.Template.Name,
		TemplateVersion:         svc.Edges.Template.Version,
		WorkflowStepExecutionID: opts.WorkflowStepExecutionID,
		Attributes:              svc.Attributes,
		DeployerType:            DeployerType,
	}

	status.ServiceRevisionStatusRunning.Unknown(entity, "")
	entity.Status.SetSummary(status.WalkServiceRevision(&entity.Status))

	// Inherit the output of previous revision to create a new one.
	if prevEntity != nil {
		entity.Output = prevEntity.Output
	}

	switch {
	case opts.JobType == JobTypeApply && entity.Output != "":
		// Get required providers from the previous output after first deployment.
		requiredProviders, err := m.getRequiredProviders(ctx, opts.ServiceID, entity.Output)
		if err != nil {
			return nil, err
		}
		entity.PreviousRequiredProviders = requiredProviders
	case opts.JobType == JobTypeDestroy && entity.Output != "":
		if status.ServiceRevisionStatusReady.IsFalse(prevEntity) {
			// Get required providers from the previous output after first deployment.
			requiredProviders, err := m.getRequiredProviders(ctx, opts.ServiceID, entity.Output)
			if err != nil {
				return nil, err
			}
			entity.PreviousRequiredProviders = requiredProviders
		} else {
			// Copy required providers from the previous revision.
			entity.PreviousRequiredProviders = prevEntity.PreviousRequiredProviders
			// Reuse other fields from the previous revision.
			entity.TemplateName = prevEntity.TemplateName
			entity.TemplateVersion = prevEntity.TemplateVersion
			entity.Attributes = prevEntity.Attributes
			entity.InputPlanConfigs = prevEntity.InputPlanConfigs
		}
	}

	// Create revision.
	entity, err = m.modelClient.ServiceRevisions().Create().
		Set(entity).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m Manager) Update(
	ctx context.Context,
	revision *model.ServiceRevision,
) error {
	revision, err := m.modelClient.ServiceRevisions().UpdateOne(revision).
		Set(revision).
		Save(ctx)
	if err != nil {
		return err
	}

	return revisionbus.Notify(ctx, m.modelClient, revision)
}

// getRequiredProviders get required providers of the service.
func (m Manager) getRequiredProviders(
	ctx context.Context,
	serviceID object.ID,
	previousOutput string,
) ([]types.ProviderRequirement, error) {
	stateRequiredProviderSet := sets.NewString()

	previousRequiredProviders, err := m.getPreviousRequiredProviders(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	stateRequiredProviders, err := parser.ParseStateProviders(previousOutput)
	if err != nil {
		return nil, err
	}

	stateRequiredProviderSet.Insert(stateRequiredProviders...)

	requiredProviders := make([]types.ProviderRequirement, 0, len(previousRequiredProviders))

	for _, p := range previousRequiredProviders {
		if stateRequiredProviderSet.Has(p.Name) {
			requiredProviders = append(requiredProviders, p)
		}
	}

	return requiredProviders, nil
}

// getPreviousRequiredProviders get previous succeed revision required providers.
// NB(alex): the previous revision may be failed, the failed revision may not contain required providers of states.
func (m Manager) getPreviousRequiredProviders(
	ctx context.Context,
	serviceID object.ID,
) ([]types.ProviderRequirement, error) {
	prevRequiredProviders := make([]types.ProviderRequirement, 0)

	entity, err := m.modelClient.ServiceRevisions().Query().
		Where(servicerevision.ServiceID(serviceID)).
		Order(model.Desc(servicerevision.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if entity == nil {
		return prevRequiredProviders, nil
	}

	templateVersion, err := m.modelClient.TemplateVersions().Query().
		Where(
			templateversion.Name(entity.TemplateName),
			templateversion.Version(entity.TemplateVersion),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if templateVersion.Schema != nil {
		prevRequiredProviders = append(prevRequiredProviders, templateVersion.Schema.RequiredProviders...)
	}

	return prevRequiredProviders, nil
}
