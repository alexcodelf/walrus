package operation

import (
	"context"

	core "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"

	wf "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/resourceruns/deployer"
	"github.com/seal-io/walrus/pkg/resourceruns/deployers"
	"github.com/seal-io/walrus/pkg/resourceruns/kubehelper"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemauthz"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	WorkflowTemplateEntryPoint = "entrypoint"
)

type ResourceRunWorkflowManager struct {
	client clientset.Interface
}

func NewResourceRunWorkflowManager() *ResourceRunWorkflowManager {
	client := system.LoopbackKubeClient.Get()
	return &ResourceRunWorkflowManager{
		client: client,
	}
}

func (w *ResourceRunWorkflowManager) BuildPlanWorkflow(
	ctx context.Context,
	client clientset.Interface,
	run *walruscore.ResourceRun,
) (*wf.Workflow, error) {
	workflow := &wf.Workflow{
		ObjectMeta: meta.ObjectMeta{
			Name:      kubehelper.NormalizeResourceRunPlanWorkflowName(run.Name),
			Namespace: run.Namespace,
			OwnerReferences: []meta.OwnerReference{
				{
					APIVersion: walruscore.SchemeGroupVersion.String(),
					Kind:       "ResourceRun",
					Name:       run.Name,
					UID:        run.UID,
				},
			},
			Annotations: map[string]string{
				"walrus.io/resource-run": run.Name,
				"walrus.io/environment":  run.Namespace,
			},
			Labels: w.GetWorkflowDefaultLabels(),
		},
		Spec: wf.WorkflowSpec{
			Entrypoint: WorkflowTemplateEntryPoint,
		},
	}

	res, err := client.WalruscoreV1().Resources(run.Namespace).Get(ctx, run.Spec.ResourceName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	runTemplate, err := client.WalruscoreV1().ResourceRunTemplates(system.NamespaceName).Get(
		ctx, run.Status.ResourceRunTemplate.Name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	resHook, err := client.WalruscoreV1().ResourceHooks(res.Status.ResourceHook.Namespace).Get(
		ctx, res.Status.ResourceHook.Name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	planSteps, err := ComputeResourcePlanSteps(ctx, w.client, runTemplate, resHook)
	if err != nil {
		return nil, err
	}

	templates, err := w.buildWorkflowTemplates(ctx, run, planSteps)
	if err != nil {
		return nil, err
	}
	workflow.Spec.Templates = templates

	return workflow, nil
}

func (w *ResourceRunWorkflowManager) CreatePlanWorkflow(
	ctx context.Context,
	run *walruscore.ResourceRun,
) error {
	planWorkflow, err := w.BuildPlanWorkflow(ctx, w.client, run)
	if err != nil {
		return err
	}

	sa, err := w.CreateResourceRunWorkflowServiceAccount(ctx, run)
	if err != nil {
		return err
	}
	planWorkflow.Spec.ServiceAccountName = sa.Name

	_, err = w.client.ArgoprojworkflowV1alpha1().Workflows(planWorkflow.Namespace).Create(ctx, planWorkflow, meta.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (w *ResourceRunWorkflowManager) CreateApproveApplyWorkflow(
	ctx context.Context,
	run *walruscore.ResourceRun,
) error {
	approveApplyWorkflow, err := w.BuildApproveApplyWorkflow(ctx, w.client, run)
	if err != nil {
		return err
	}

	sa, err := w.CreateResourceRunWorkflowServiceAccount(ctx, run)
	if err != nil {
		return err
	}
	approveApplyWorkflow.Spec.ServiceAccountName = sa.Name

	_, err = w.client.ArgoprojworkflowV1alpha1().Workflows(approveApplyWorkflow.Namespace).Create(ctx, approveApplyWorkflow, meta.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (w *ResourceRunWorkflowManager) BuildApproveApplyWorkflow(
	ctx context.Context,
	client clientset.Interface,
	run *walruscore.ResourceRun,
) (*wf.Workflow, error) {
	workflow := &wf.Workflow{
		ObjectMeta: meta.ObjectMeta{
			Name:      kubehelper.NormalizeResourceRunApproveApplyWorkflowName(run.Name),
			Namespace: run.Namespace,
			OwnerReferences: []meta.OwnerReference{
				{
					APIVersion: walruscore.SchemeGroupVersion.String(),
					Kind:       "ResourceRun",
					Name:       run.Name,
					UID:        run.UID,
				},
			},
			Annotations: map[string]string{
				"walrus.io/resource-run": run.Name,
				"walrus.io/environment":  run.Namespace,
			},
			Labels: w.GetWorkflowDefaultLabels(),
		},
		Spec: wf.WorkflowSpec{
			Entrypoint: WorkflowTemplateEntryPoint,
		},
	}

	res, err := client.WalruscoreV1().Resources(run.Namespace).Get(ctx, run.Spec.ResourceName, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	runTemplate, err := client.WalruscoreV1().ResourceRunTemplates(system.NamespaceName).Get(
		ctx, run.Status.ResourceRunTemplate.Name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	resHook, err := client.WalruscoreV1().ResourceHooks(res.Status.ResourceHook.Namespace).Get(
		ctx, res.Status.ResourceHook.Name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	approveApplySteps, err := ComputeResourceApproveApplySteps(ctx, w.client, runTemplate, resHook)
	if err != nil {
		return nil, err
	}

	templates, err := w.buildWorkflowTemplates(ctx, run, approveApplySteps)
	if err != nil {
		return nil, err
	}
	workflow.Spec.Templates = templates

	return workflow, nil
}

func (w *ResourceRunWorkflowManager) buildWorkflowTemplates(
	ctx context.Context,
	run *walruscore.ResourceRun,
	steps []walruscore.ResourceRunStep,
) ([]wf.Template, error) {
	wfTemplates := make([]wf.Template, 0)

	for _, step := range steps {
		var wfTemplate wf.Template
		switch step.Type {
		case walruscore.ResourceRunStepTypeContainer:
			wfTemplate = wf.Template{
				Name:      step.Name,
				Container: step.Container,
			}

		case walruscore.ResourceRunStepTypeApproval:
			wfTemplate = wf.Template{
				Name:    step.Name,
				Suspend: &wf.SuspendTemplate{},
			}
		case walruscore.ResourceRunStepTypePlan:
			template, err := w.client.WalrusV1().Templates(run.Spec.Template.Namespace).Get(ctx, run.Spec.Template.Name, meta.GetOptions{})
			if err != nil {
				return nil, err
			}

			dp, err := deployers.NewDeployer(deployer.CreateOptions{
				Type: template.Spec.TemplateFormat,
			})
			if err != nil {
				return nil, err
			}

			wfTemplatePtr, err := dp.Plan(ctx, run, deployer.PlanOptions{})
			if err != nil {
				return nil, err
			}

			wfTemplate = *wfTemplatePtr

		case walruscore.ResourceRunStepTypeApply:
			template, err := w.client.WalrusV1().Templates(run.Spec.Template.Namespace).Get(ctx, run.Spec.Template.Name, meta.GetOptions{})
			if err != nil {
				return nil, err
			}

			dp, err := deployers.NewDeployer(deployer.CreateOptions{
				Type: template.Spec.TemplateFormat,
			})
			if err != nil {
				return nil, err
			}

			wfTemplatePtr, err := dp.Apply(ctx, run, deployer.ApplyOptions{})
			if err != nil {
				return nil, err
			}

			wfTemplate = *wfTemplatePtr
		}

		wfTemplates = append(wfTemplates, wfTemplate)
	}

	entrypointSteps := make([]wf.WorkflowStep, 0)

	for i := range wfTemplates {
		entrypointSteps = append(entrypointSteps, wf.WorkflowStep{
			Name:     wfTemplates[i].Name,
			Template: wfTemplates[i].Name,
		})
	}

	templateEntrypoint := wf.Template{
		Name: WorkflowTemplateEntryPoint,
		Steps: []wf.ParallelSteps{
			{
				Steps: entrypointSteps,
			},
		},
	}

	wfTemplates = append(wfTemplates, templateEntrypoint)

	return wfTemplates, nil
}

func (w *ResourceRunWorkflowManager) GetWorkflowDefaultLabels() map[string]string {
	labels := make(map[string]string)

	disableApps := system.DisableApplications.Get()
	if !disableApps.Has("*") && !disableApps.Has("argo-workflow") {
		labels["workflows.argoproj.io/controller-instanceid"] = "walrus-workflows"
	}

	return labels
}

func (w *ResourceRunWorkflowManager) CreateResourceRunWorkflowServiceAccount(
	ctx context.Context,
	run *walruscore.ResourceRun,
) (*core.ServiceAccount, error) {
	// If the service account already exists, return it.
	sa, err := w.client.CoreV1().ServiceAccounts(run.Namespace).Get(ctx, kubehelper.NormalizeResourceRunServiceAccountName(run.Name), meta.GetOptions{})
	if err != nil && !kerrors.IsNotFound(err) {
		return nil, err
	}
	if sa != nil && sa.Name != "" {
		return sa, nil
	}

	// Create service account for the template.
	ownerReference := meta.OwnerReference{
		APIVersion: walruscore.SchemeGroupVersion.String(),
		Kind:       "ResourceRun",
		Name:       run.Name,
		UID:        run.UID,
	}

	sa = &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Namespace:       run.Namespace,
			Name:            kubehelper.NormalizeResourceRunServiceAccountName(run.Name),
			OwnerReferences: []meta.OwnerReference{ownerReference},
		},
	}

	sa, err = w.client.CoreV1().ServiceAccounts(sa.Namespace).Create(ctx, sa, meta.CreateOptions{})
	if err != nil {
		return nil, err
	}

	rb := &rbac.RoleBinding{
		ObjectMeta: meta.ObjectMeta{
			Namespace:       run.Namespace,
			Name:            kubehelper.NormalizeResourceRunRoleBindingName(run.Name),
			OwnerReferences: []meta.OwnerReference{ownerReference},
		},
		RoleRef: rbac.RoleRef{
			APIGroup: rbac.GroupName,
			Kind:     "ClusterRole",
			Name:     systemauthz.DeployerClusterRoleName,
		},
		Subjects: []rbac.Subject{
			{
				Namespace: sa.Namespace,
				Kind:      "ServiceAccount",
				Name:      sa.Name,
			},
		},
	}

	_, err = w.client.RbacV1().RoleBindings(sa.Namespace).Create(ctx, rb, meta.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return sa, nil
}
