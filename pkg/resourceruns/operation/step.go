package operation

import (
	"context"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func computeResourcePlanStep() walruscore.ResourceRunStep {
	return walruscore.ResourceRunStep{
		Name: "plan",
		Type: walruscore.ResourceRunStepTypePlan,
	}
}

func computeResourceRunApproveStep() walruscore.ResourceRunStep {
	return walruscore.ResourceRunStep{
		Name: "approve",
		Type: walruscore.ResourceRunStepTypeApproval,
		Template: &walruscore.ResourceRunStepTemplateReference{
			Namespace: system.NamespaceName,
			Name:      systemkuberes.ResourceRunStepTemplateDefaultApproval,
		},
	}
}

func computeResourceRunApplyStep() walruscore.ResourceRunStep {
	return walruscore.ResourceRunStep{
		Name: "apply",
		Type: walruscore.ResourceRunStepTypeApply,
	}
}

func ComputeResourcePlanSteps(
	ctx context.Context,
	client clientset.Interface,
	runTemplate *walruscore.ResourceRunTemplate,
	resHook *walruscore.ResourceHook,
) ([]walruscore.ResourceRunStep, error) {
	prePlanHooks := mergeHooks(runTemplate.Spec.Plan.Pre, resHook.Spec.Plan.Pre)
	postPlanHooks := mergeHooks(resHook.Spec.Plan.Post, runTemplate.Spec.Plan.Post)

	// Get resource run pre-plan steps and post-plan steps.
	preSteps, err := convertHooks(ctx, client, prePlanHooks)
	if err != nil {
		return nil, err
	}

	postSteps, err := convertHooks(ctx, client, postPlanHooks)
	if err != nil {
		return nil, err
	}

	steps := make([]walruscore.ResourceRunStep, 0)
	// Add pre-plan steps.
	steps = append(steps, preSteps...)

	// Add plan step.
	planStep := computeResourcePlanStep()
	steps = append(steps, planStep)

	// Add post-plan steps.
	steps = append(steps, postSteps...)

	return steps, nil
}

func ComputeResourceApproveApplySteps(
	ctx context.Context,
	client clientset.Interface,
	runTemplate *walruscore.ResourceRunTemplate,
	resHook *walruscore.ResourceHook,
) ([]walruscore.ResourceRunStep, error) {
	preApproveHooks := mergeHooks(runTemplate.Spec.Approve.Pre, resHook.Spec.Approve.Pre)
	postApproveHooks := mergeHooks(resHook.Spec.Approve.Post, runTemplate.Spec.Approve.Post)

	preApplyHooks := mergeHooks(runTemplate.Spec.Apply.Pre, resHook.Spec.Apply.Pre)
	postApplyHooks := mergeHooks(resHook.Spec.Apply.Post, runTemplate.Spec.Apply.Post)

	steps := make([]walruscore.ResourceRunStep, 0)

	// Handler pre-approve and post-approve steps.
	{
		preApproveSteps, err := convertHooks(ctx, client, preApproveHooks)
		if err != nil {
			return nil, err
		}

		steps = append(steps, preApproveSteps...)
		// Add approve step.
		steps = append(steps, computeResourceRunApproveStep())

		postApproveSteps, err := convertHooks(ctx, client, postApproveHooks)
		if err != nil {
			return nil, err
		}

		steps = append(steps, postApproveSteps...)
	}

	// Handler pre-apply and post-apply steps.
	{
		preApplySteps, err := convertHooks(ctx, client, preApplyHooks)
		if err != nil {
			return nil, err
		}

		steps = append(steps, preApplySteps...)
		// Add apply step.
		steps = append(steps, computeResourceRunApplyStep())

		postApplySteps, err := convertHooks(ctx, client, postApplyHooks)
		if err != nil {
			return nil, err
		}

		steps = append(steps, postApplySteps...)
	}

	return steps, nil
}

// ComputeResourceRunSteps computes the steps of the resource run.
// It will compute the steps of the resource run based on the resource run template and resource hook.
// The steps will be computed in the following order:
// Plan -> Approve -> Apply(with their pre and post steps).
func ComputeResourceRunSteps(ctx context.Context, res *walruscore.Resource, run *walruscore.ResourceRun) ([]walruscore.ResourceRunStep, error) {
	client := system.LoopbackKubeClient.Get()

	runTemplate, err := client.WalruscoreV1().ResourceRunTemplates(system.NamespaceName).Get(
		ctx, run.Status.ResourceRunTemplate.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	resHook, err := client.WalruscoreV1().ResourceHooks(res.Status.ResourceHook.Namespace).Get(
		ctx, res.Status.ResourceHook.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	planSteps, err := ComputeResourcePlanSteps(ctx, client, runTemplate, resHook)
	if err != nil {
		return nil, err
	}

	approveApplySteps, err := ComputeResourceApproveApplySteps(ctx, client, runTemplate, resHook)
	if err != nil {
		return nil, err
	}

	steps := make([]walruscore.ResourceRunStep, 0)
	steps = append(steps, planSteps...)
	steps = append(steps, approveApplySteps...)

	return steps, nil
}

func setStepType(ctx context.Context, client clientset.Interface, step *walruscore.ResourceRunStep) error {
	if step.Template != nil {
		template, err := client.WalruscoreV1().ResourceRunStepTemplates(step.Template.Namespace).Get(
			ctx, step.Template.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}

		switch {
		case template.Spec.Approval != nil:
			step.Type = walruscore.ResourceRunStepTypeApproval
		case template.Spec.Container != nil:
			step.Type = walruscore.ResourceRunStepTypeContainer
		}
	}

	if step.Container != nil {
		step.Type = walruscore.ResourceRunStepTypeContainer
	}

	return fmt.Errorf("invalid step type")
}

func convertHooks(ctx context.Context, client clientset.Interface, hooks []walruscore.ResourceHookStep) ([]walruscore.ResourceRunStep, error) {
	var hookSteps []walruscore.ResourceRunStep

	for _, hook := range hooks {
		step := walruscore.ResourceRunStep{
			Name:      hook.Name,
			Template:  hook.ResourceRunStepTemplate,
			Container: hook.Container,
		}

		if err := setStepType(ctx, client, &step); err != nil {
			return nil, err
		}

		hookSteps = append(hookSteps, step)
	}

	return hookSteps, nil
}

func mergeHooks(steps, mergeSteps []walruscore.ResourceHookStep) []walruscore.ResourceHookStep {
	hooks := make([]walruscore.ResourceHookStep, 0)

	hooks = append(hooks, steps...)
	hooks = append(hooks, mergeSteps...)

	return hooks
}
