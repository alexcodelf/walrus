package resourcerun

import (
	"context"
	"errors"
	"fmt"

	busrun "github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/resourcedefinitions"
	"github.com/seal-io/walrus/pkg/terraform/parser"
	"github.com/seal-io/walrus/utils/log"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	// DeployerType is the type of the deployer.
	DeployerType = types.DeployerTypeTF

	JobTypeApply   = "apply"
	JobTypeDestroy = "destroy"
)

// Manager is the manager of resource run.
// It is responsible for creating resource run and loading input configs.
// Model client should be passed in functions to avoid transaction issue.
type Manager struct {
	logger log.Logger

	InputLoader InputLoader
}

func NewManager() *Manager {
	return &Manager{
		InputLoader: NewInputLoader(types.DeployerTypeTF),
	}
}

// CreateOptions are the options for creating a resource run.
type CreateOptions struct {
	// ResourceID indicates the ID of resource which is for create the revision.
	ResourceID object.ID

	// JobType indicates the type of the job.
	JobType string

	// ChangeComment is the comment of the change.
	ChangeComment string
}

func (m Manager) Create(ctx context.Context, mc model.ClientSet, opts CreateOptions) (*model.ResourceRun, error) {
	// Validate if there is a running run.
	prevEntity, err := mc.ResourceRuns().Query().
		Where(resourcerun.And(
			resourcerun.ResourceID(opts.ResourceID),
			resourcerun.DeployerType(DeployerType))).
		Order(model.Desc(resourcerun.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if prevEntity != nil && status.ResourceRunStatusReady.IsUnknown(prevEntity) {
		return nil, errors.New("deployment is running")
	}

	// Get the corresponding resource and template version.
	res, err := mc.Resources().Query().
		Where(resource.ID(opts.ResourceID)).
		WithTemplate(func(tvq *model.TemplateVersionQuery) {
			tvq.Select(
				templateversion.FieldName,
				templateversion.FieldVersion,
				templateversion.FieldTemplateID)
		}).
		WithProject(func(pq *model.ProjectQuery) {
			pq.Select(project.FieldName)
		}).
		WithEnvironment(func(env *model.EnvironmentQuery) {
			env.Select(environment.FieldLabels)
			env.Select(environment.FieldName)
			env.Select(environment.FieldType)
		}).
		WithResourceDefinition(func(rd *model.ResourceDefinitionQuery) {
			rd.Select(resourcedefinition.FieldType)
			rd.WithMatchingRules(func(mrq *model.ResourceDefinitionMatchingRuleQuery) {
				mrq.Order(model.Asc(resourcedefinitionmatchingrule.FieldOrder)).
					Select(
						resourcedefinitionmatchingrule.FieldName,
						resourcedefinitionmatchingrule.FieldSelector,
						resourcedefinitionmatchingrule.FieldAttributes,
					).
					WithTemplate(func(tvq *model.TemplateVersionQuery) {
						tvq.Select(
							templateversion.FieldID,
							templateversion.FieldVersion,
							templateversion.FieldName,
						)
					})
			})
		}).
		Select(
			resource.FieldID,
			resource.FieldProjectID,
			resource.FieldEnvironmentID,
			resource.FieldType,
			resource.FieldLabels,
			resource.FieldAnnotations,
			resource.FieldAttributes).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var (
		templateID                    object.ID
		templateName, templateVersion string
		attributes                    property.Values
	)

	switch {
	case res.TemplateID != nil:
		templateID = res.Edges.Template.TemplateID
		templateName = res.Edges.Template.Name
		templateVersion = res.Edges.Template.Version
		attributes = res.Attributes
	case res.ResourceDefinitionID != nil:
		rd := res.Edges.ResourceDefinition
		matchRule := resourcedefinitions.Match(
			rd.Edges.MatchingRules,
			res.Edges.Project.Name,
			res.Edges.Environment.Name,
			res.Edges.Environment.Type,
			res.Edges.Environment.Labels,
			res.Labels,
		)

		if matchRule == nil {
			return nil, fmt.Errorf("resource definition %s does not match resource %s", rd.Name, res.Name)
		}

		_, err = mc.Resources().UpdateOne(res).
			SetResourceDefinitionMatchingRuleID(matchRule.ID).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		templateName = matchRule.Edges.Template.Name
		templateVersion = matchRule.Edges.Template.Version

		templateID, err = mc.Templates().Query().
			Where(
				template.Name(templateName),
				// Now we only support resource definition globally.
				template.ProjectIDIsNil(),
			).
			OnlyID(ctx)
		if err != nil {
			return nil, err
		}

		// Merge attributes. Resource attributes take precedence over resource definition attributes.
		attributes = matchRule.Attributes
		if attributes == nil {
			attributes = make(property.Values)
		}

		for k, v := range res.Attributes {
			attributes[k] = v
		}
	default:
		return nil, errors.New("missing template or resource definition")
	}

	userSubject, err := getSubject(ctx, mc, res)
	if err != nil {
		return nil, err
	}

	entity := &model.ResourceRun{
		ProjectID:       res.ProjectID,
		EnvironmentID:   res.EnvironmentID,
		ResourceID:      res.ID,
		TemplateID:      templateID,
		TemplateName:    templateName,
		TemplateVersion: templateVersion,
		Attributes:      attributes,
		DeployerType:    DeployerType,
		CreatedBy:       userSubject.Name,
		ChangeComment:   opts.ChangeComment,
	}

	status.ResourceRunStatusReady.Unknown(entity, "")
	entity.Status.SetSummary(status.WalkResourceRun(&entity.Status))

	// Inherit the output of previous run to create a new one.
	if prevEntity != nil {
		entity.Output = prevEntity.Output
	}

	switch {
	case opts.JobType == JobTypeApply && entity.Output != "":
		// Get required providers from the previous output after first deployment.
		requiredProviders, err := m.getRequiredProviders(ctx, mc, opts.ResourceID, entity.Output)
		if err != nil {
			return nil, err
		}
		entity.PreviousRequiredProviders = requiredProviders
	case opts.JobType == JobTypeDestroy && entity.Output != "":
		if status.ResourceRunStatusReady.IsFalse(prevEntity) {
			// Get required providers from the previous output after first deployment.
			requiredProviders, err := m.getRequiredProviders(ctx, mc, opts.ResourceID, entity.Output)
			if err != nil {
				return nil, err
			}
			entity.PreviousRequiredProviders = requiredProviders
		} else {
			// Copy required providers from the previous run.
			entity.PreviousRequiredProviders = prevEntity.PreviousRequiredProviders
			// Reuse other fields from the previous run.
			entity.TemplateID = prevEntity.TemplateID
			entity.TemplateName = prevEntity.TemplateName
			entity.TemplateVersion = prevEntity.TemplateVersion
			entity.Attributes = prevEntity.Attributes
			entity.InputPlan = prevEntity.InputPlan
		}
	}

	// Create run.
	entity, err = mc.ResourceRuns().Create().
		Set(entity).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m Manager) UpdateRunStatus(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.ResourceRun,
	run *model.ResourceRun,
) error {
	// Report to resource run.
	run.Status.SetSummary(status.WalkResourceRun(&run.Status))

	run, err := mc.ResourceRuns().UpdateOne(run).
		SetStatus(run.Status).
		Save(ctx)
	if err != nil {
		return err
	}

	if err = busrun.Notify(ctx, mc, run); err != nil {
		m.logger.Error(err)
		return err
	}

	return nil
}

// getRequiredProviders gets the required providers from the output of the previous run.
func (m Manager) getRequiredProviders(
	ctx context.Context,
	mc model.ClientSet,
	resourceID object.ID,
	previousOutput string,
) ([]types.ProviderRequirement, error) {
	stateRequiredProviderSet := sets.NewString()

	previousRequiredProviders, err := m.getPreviousRequiredProviders(ctx, mc, resourceID)
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

// getPreviousRequiredProviders get previous succeed run required providers.
// NB(alex): the previous run may be failed, the failed run may not contain required providers of states.
func (m Manager) getPreviousRequiredProviders(
	ctx context.Context,
	mc model.ClientSet,
	resourceID object.ID,
) ([]types.ProviderRequirement, error) {
	prevRequiredProviders := make([]types.ProviderRequirement, 0)

	entity, err := mc.ResourceRuns().Query().
		Where(resourcerun.ResourceID(resourceID)).
		Order(model.Desc(resourcerun.FieldCreateTime)).
		First(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if entity == nil {
		return prevRequiredProviders, nil
	}

	templateVersion, err := mc.TemplateVersions().Query().
		Where(
			templateversion.TemplateID(entity.TemplateID),
			templateversion.Version(entity.TemplateVersion),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if len(templateVersion.Schema.RequiredProviders) != 0 {
		prevRequiredProviders = append(prevRequiredProviders, templateVersion.Schema.RequiredProviders...)
	}

	return prevRequiredProviders, nil
}

func (m Manager) LoadInputConfigs(
	ctx context.Context,
	mc model.ClientSet,
	opts *ConfigLoaderOptions,
) (map[string]ConfigData, error) {
	return m.InputLoader.LoadAll(ctx, mc, opts)
}

func (m Manager) LoadProviderConfigs(conns model.Connectors) (map[string]ConfigData, error) {
	return m.InputLoader.LoadProviders(conns)
}
