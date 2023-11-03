package workflow

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"

	apiservice "github.com/seal-io/walrus/pkg/apis/service"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

func init() {
	stepValidatorCreators = map[string]func(*model.WorkflowStepCreateInput) WorkflowStepValidator{
		types.WorkflowStepTypeService: func(step *model.WorkflowStepCreateInput) WorkflowStepValidator {
			return &WorkflowStepServiceValidator{step}
		},
		types.WorkflowStepTypeApproval: func(step *model.WorkflowStepCreateInput) WorkflowStepValidator {
			return &WorkflowStepApprovalValidator{step}
		},
	}
}

func validateStages(ctx *gin.Context, client *model.Client, stages []*model.WorkflowStageCreateInput) error {
	for _, stage := range stages {
		if err := validateStage(ctx, client, stage); err != nil {
			return err
		}
	}

	return nil
}

func validateStage(ctx *gin.Context, client *model.Client, stage *model.WorkflowStageCreateInput) error {
	if err := validateSteps(ctx, client, stage.Steps); err != nil {
		return fmt.Errorf("invalid steps: %w", err)
	}

	return nil
}

func validateSteps(ctx *gin.Context, client *model.Client, steps []*model.WorkflowStepCreateInput) error {
	for _, step := range steps {
		if err := validateStep(ctx, client, step); err != nil {
			return err
		}
	}

	return nil
}

func validateStep(ctx *gin.Context, client *model.Client, step *model.WorkflowStepCreateInput) error {
	creator, ok := stepValidatorCreators[step.Type]
	if !ok {
		return fmt.Errorf("unknown step type: %s", step.Type)
	}

	err := checkSpec(step.Spec)
	if err != nil {
		return fmt.Errorf("invalid spec: %w", err)
	}

	stepValidator := creator(step)

	return stepValidator.Validate(ctx, client)
}

type WorkflowStepValidator interface {
	Set(*model.WorkflowStepCreateInput)
	// Validate validates the spec of the workflow step.
	Validate(*gin.Context, *model.Client) error
}

// WorkflowStepServiceValidator validates the spec of a service workflow step.
type WorkflowStepServiceValidator struct {
	*model.WorkflowStepCreateInput `path:",inline" json:",inline"`
}

type ServiceCreateMeta struct {
	Project     *model.ProjectQueryInput     `json:"project"`
	Environment *model.EnvironmentQueryInput `json:"environment"`
}

func (s *WorkflowStepServiceValidator) Set(input *model.WorkflowStepCreateInput) {
	s.WorkflowStepCreateInput = input
}

func (s *WorkflowStepServiceValidator) Validate(ctx *gin.Context, client *model.Client) error {
	sci := &model.ServiceCreateInput{}
	scm := &ServiceCreateMeta{}

	sci.SetGinContext(ctx)
	sci.SetModelClient(client)

	jsonData, err := json.Marshal(s.Spec)
	if err != nil {
		return fmt.Errorf("failed to marshal service spec: %w", err)
	}

	if err := json.Unmarshal(jsonData, sci); err != nil {
		return fmt.Errorf("failed to unmarshal service input: %w", err)
	}

	if err := json.Unmarshal(jsonData, scm); err != nil {
		return fmt.Errorf("failed to unmarshal service meta: %w", err)
	}

	sci.Project = scm.Project
	sci.Environment = scm.Environment

	if err := apiservice.ValidateCreateInput(*sci); err != nil {
		return err
	}

	return nil
}

// WorkflowStepApprovalValidator validates the spec of an approval workflow step.
type WorkflowStepApprovalValidator struct {
	*model.WorkflowStepCreateInput
}

func (s *WorkflowStepApprovalValidator) Set(input *model.WorkflowStepCreateInput) {
	s.WorkflowStepCreateInput = input
}

func (s *WorkflowStepApprovalValidator) Validate(*gin.Context, *model.Client) error {
	approvalType, ok := s.Spec[types.WorkflowStepApprovalType].(string)
	if !ok {
		return fmt.Errorf("invalid approval type")
	}

	switch approvalType {
	case types.WorkflowStepApprovalTypeOr, types.WorkflowStepApprovalTypeAnd:
	default:
		return fmt.Errorf("invalid approval type: %s", approvalType)
	}

	approvalUsers, ok := s.Spec[types.WorkflowStepApprovalUsers].([]object.ID)
	if !ok {
		return fmt.Errorf("invalid approval users")
	}

	s.Spec = map[string]any{
		types.WorkflowStepApprovalType:  approvalType,
		types.WorkflowStepApprovalUsers: approvalUsers,
	}

	return nil
}

var stepValidatorCreators = map[string]func(*model.WorkflowStepCreateInput) WorkflowStepValidator{}

// checkSpec check the spec of workflow step contains any {{xxx}}.
func checkSpec(spec map[string]any) error {
	bs, err := json.Marshal(spec)
	if err != nil {
		return err
	}

	// Argo workflow template use {{xxx}} as template keyword.
	// Make sure the spec does not contain any {{xxx}}.
	keywordsReg := regexp.MustCompile(`{{.*}}`)
	matches := keywordsReg.FindAll(bs, -1)

	if len(matches) > 0 {
		return fmt.Errorf("spec contains keywords: %s", string(matches[0]))
	}

	return nil
}
