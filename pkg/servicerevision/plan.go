package servicerevision

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	pkgenv "github.com/seal-io/walrus/pkg/environment"
	pkgservice "github.com/seal-io/walrus/pkg/service"
)

type IPlan interface {
	// Plan plans the revision.
	LoadPlan(context.Context, PlanOptions) ([]byte, error)
	// LoadConfigs loads the plan configs of the plan options.
	LoadConfigs(context.Context, PlanOptions) (map[string][]byte, error)
	// LoadConnectorConfigs loads the connector configs of the plan options.
	// Some connectors may be required to deploy the service.
	LoadConnectorConfigs(model.Connectors) (map[string][]byte, error)
}

// PlanOptions are the options for planning a revision.
type PlanOptions struct {
	// SecretMountPath of the deploy job.
	SecretMountPath string

	ServiceRevision *model.ServiceRevision
	Connectors      model.Connectors
	ProjectID       object.ID
	EnvironmentID   object.ID
	SubjectID       object.ID
	// Metadata.
	ProjectName          string
	EnvironmentName      string
	ServiceName          string
	ServiceID            object.ID
	ManagedNamespaceName string
}

// NewPlan creates a new plan with the plan type.
func NewPlan(planType string, mc model.ClientSet) IPlan {
	switch planType {
	case types.DeployerTypeTF:
		return NewTerraformPlan(mc)
	default:
		return nil
	}
}

// SetPlanOptions sets the plan options.
func SetPlanOptions(ctx context.Context, mc model.ClientSet, opts *PlanOptions) error {
	if !status.ServiceRevisionStatusRunning.IsUnknown(opts.ServiceRevision) {
		return errors.New("service revision is not running")
	}

	connectors, err := pkgenv.GetConnectors(ctx, mc, opts.ServiceRevision.EnvironmentID)
	if err != nil {
		return err
	}

	proj, err := mc.Projects().Get(ctx, opts.ServiceRevision.ProjectID)
	if err != nil {
		return err
	}

	env, err := dao.GetEnvironmentByID(ctx, mc, opts.ServiceRevision.EnvironmentID)
	if err != nil {
		return err
	}

	svc, err := mc.Services().Get(ctx, opts.ServiceRevision.ServiceID)
	if err != nil {
		return err
	}

	var subjectID object.ID

	sj, _ := session.GetSubject(ctx)
	if sj.ID != "" {
		subjectID = sj.ID
	} else {
		subjectID, err = pkgservice.GetSubjectID(svc)
		if err != nil {
			return err
		}
	}

	if subjectID == "" {
		return errors.New("subject id is empty")
	}

	opts.Connectors = connectors
	opts.ProjectID = proj.ID
	opts.EnvironmentID = env.ID
	opts.SubjectID = subjectID

	// Metadata.
	opts.ProjectName = proj.Name
	opts.EnvironmentName = env.Name
	opts.ServiceName = svc.Name
	opts.ServiceID = svc.ID
	opts.ManagedNamespaceName = pkgenv.GetManagedNamespaceName(env)

	return nil
}
