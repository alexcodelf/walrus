package operation

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/kubemeta"
	reskubehelper "github.com/seal-io/walrus/pkg/resources/kubehelper"
	"github.com/seal-io/walrus/pkg/system"
)

const (
	ResourceRunTemplateDefault = "default"
)

func CreateResourceRun(ctx context.Context, res *walruscore.Resource, runType walruscore.ResourceRunType) (*walruscore.ResourceRun, error) {
	switch runType {
	case walruscore.ResourceRunTypeCreate,
		walruscore.ResourceRunTypeUpdate,
		walruscore.ResourceRunTypeDelete,
		walruscore.ResourceRunTypeStart,
		walruscore.ResourceRunTypeStop,
		walruscore.ResourceRunTypeRollback:
	default:
		panic("unsupported resource run type")
	}

	loopbackKubeClient := system.LoopbackKubeClient.Get()
	run := &walruscore.ResourceRun{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: res.Namespace,
			Name:      reskubehelper.NormalizeResourceRunName(res.Name),
		},
		Spec: walruscore.ResourceRunSpec{
			Project:      res.Status.Phase,
			ResourceName: res.Name,
			Type:         runType,
			Attributes:   res.Spec.Attributes,
			Template:     *res.Spec.Template,
		},
	}

	kubemeta.ControlOn(run, res, walrus.SchemeGroupVersion.WithKind("Resource"))

	return loopbackKubeClient.WalruscoreV1().ResourceRuns(res.Namespace).Create(ctx, run, metav1.CreateOptions{})
}
