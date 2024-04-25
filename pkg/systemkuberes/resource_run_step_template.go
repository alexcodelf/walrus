package systemkuberes

import (
	"context"
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceRunStepTemplateDefaultApproval = "default-approval"
)

// InstallBuiltinResourceRunStepTemplate installs the builtin resource run step template.
func InstallBuiltinResourceRunStepTemplate(ctx context.Context, cli clientset.Interface) error {
	resourceRunStepTemplate := &walruscore.ResourceRunStepTemplate{
		ObjectMeta: meta.ObjectMeta{
			Namespace: SystemNamespaceName,
			Name:      ResourceRunStepTemplateDefaultApproval,
		},
		Spec: walruscore.ResourceRunStepTemplateSpec{
			Approval: &walruscore.ResourceRunStepApprovalTemplate{
				Type: walruscore.ResourceRunStepTemplateApprovalTypeAny,
				Users: []string{
					AdminSubjectName,
				},
			},
		},
	}

	resourceRunStepTemplateCli := cli.WalruscoreV1().ResourceRunStepTemplates(SystemNamespaceName)

	_, err := kubeclientset.Create(ctx, resourceRunStepTemplateCli, resourceRunStepTemplate)
	if err != nil {
		return fmt.Errorf("install builtin resource run step template: %w", err)
	}

	return nil
}
