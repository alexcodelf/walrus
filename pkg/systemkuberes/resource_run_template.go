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
	ResourceRunTemplateDefault = "default"
)

func InstallBuiltinResourceRunTemplate(ctx context.Context, cli clientset.Interface) error {
	resourceRunTemplate := &walruscore.ResourceRunTemplate{
		ObjectMeta: meta.ObjectMeta{
			Namespace: SystemNamespaceName,
			Name:      ResourceRunTemplateDefault,
		},
		Spec: walruscore.ResourceRunTemplateSpec{},
	}

	resourceRunTemplateCli := cli.WalruscoreV1().ResourceRunTemplates(SystemNamespaceName)

	_, err := kubeclientset.Create(ctx, resourceRunTemplateCli, resourceRunTemplate)
	if err != nil {
		return fmt.Errorf("install builtin resource run template: %w", err)
	}

	return nil
}
