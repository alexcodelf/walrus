package resources

import (
	"context"
	"fmt"

	"github.com/seal-io/utils/errorx"
	kerror "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/resources/api"
	"github.com/seal-io/walrus/pkg/resources/kubehelper"
	"github.com/seal-io/walrus/pkg/system"
)

func GetSubject(obj *walruscore.Resource) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}

	subject := obj.Annotations[api.AnnotationSubject]

	return subject, nil
}

func SetSubject(ctx context.Context, resources ...*walruscore.Resource) error {
	// TODO(alex): add subject from ctx.
	sj := ""
	for i := range resources {
		if resources[i].Annotations == nil {
			resources[i].Annotations = make(map[string]string)
		}
		resources[i].Annotations[api.AnnotationSubject] = sj
	}

	return nil
}

// UpdateResourceSubject updates the subject of the resources.
func UpdateResourceSubject(ctx context.Context, resources ...*walruscore.Resource) error {
	if len(resources) == 0 {
		return nil
	}

	if err := SetSubject(ctx, resources...); err != nil {
		return err
	}

	loopbackKubeClient := system.LoopbackKubeClient.Get()

	for i := range resources {
		res := resources[i]

		_, err := loopbackKubeClient.WalruscoreV1().Resources(res.Namespace).Update(ctx, res, meta.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// IsStoppable tells whether the given resource is stoppable.
func IsStoppable(r *walruscore.Resource) bool {
	if r == nil {
		return false
	}

	if r.Labels[api.LabelResourceStoppable] == "true" ||
		(r.Spec.Template != nil && r.Labels[api.LabelResourceStoppable] != "false") {
		return true
	}

	return false
}

// CanBeStopped tells whether the given resource can be stopped.
func CanBeStopped(r *walruscore.Resource) bool {
	return apistatus.ResourceConditionDeployed.IsTrue(r)
}

func SetEnvResourceDefaultLabels(env *walrus.Environment, r *walruscore.Resource) error {
	if r == nil || env == nil {
		return errorx.Errorf("resource or environment is nil")
	}

	if r.Labels == nil {
		r.Labels = make(map[string]string)
	}

	// Only set default labels if labels stoppable are not set.
	if _, ok := r.Labels[api.LabelResourceStoppable]; ok {
		return nil
	}

	switch env.Spec.Type {
	// Dev and staging environments resources are stoppable by default.
	case walruscore.EnvironmentTypeDevelopment, walruscore.EnvironmentTypeStaging:
		r.Labels[api.LabelResourceStoppable] = "true"
	case walruscore.EnvironmentTypeProduction:
		r.Labels[api.LabelResourceStoppable] = "false"
	default:
	}

	return nil
}

// SetDefaultLabels sets default labels for the provided resources.
func SetDefaultLabels(ctx context.Context, entities ...*walruscore.Resource) error {
	if len(entities) == 0 {
		return nil
	}

	loopbackKubeClient := system.LoopbackKubeClient.Get()

	envMap := make(map[string]*walrus.Environment, len(entities))
	for _, entity := range entities {
		if _, ok := envMap[entity.Namespace]; ok {
			continue
		}

		env, err := loopbackKubeClient.WalrusV1().Environments(entity.Namespace).Get(ctx, entity.Namespace, meta.GetOptions{})
		if err != nil {
			return err
		}

		envMap[entity.Namespace] = env
	}

	for _, entity := range entities {
		env, ok := envMap[entity.Namespace]
		if !ok {
			return fmt.Errorf("environment %q not found", entity.Namespace)
		}

		if err := SetEnvResourceDefaultLabels(env, entity); err != nil {
			return err
		}
	}

	return nil
}

// GetOrCreateHook creates a resource hook if not exists.
func GetOrCreateHook(ctx context.Context, client ctrlcli.Client, res *walruscore.Resource) (*walruscore.ResourceHook, error) {
	// Create resource run hook if not exists.
	hook := &walruscore.ResourceHook{
		ObjectMeta: meta.ObjectMeta{
			Namespace: res.Namespace,
			Name:      kubehelper.NormalizeResourceHookName(res.Name),
		},
	}
	err := client.Get(ctx, ctrlcli.ObjectKeyFromObject(hook), hook)
	if err != nil {
		if !kerror.IsNotFound(err) {
			return nil, err
		}

		err = client.Create(ctx, hook)
		if err != nil {
			return nil, err
		}
	}

	return hook, nil
}
