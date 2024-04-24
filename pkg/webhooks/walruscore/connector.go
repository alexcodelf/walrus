package walruscore

import (
	"context"
	"fmt"
	"reflect"

	"github.com/seal-io/utils/stringx"
	"golang.org/x/exp/maps"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/webhook"
)

// ConnectorWebhook hooks a v1.Connector object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="connectors",scope="Namespaced",subResources=["status"]
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE","DELETE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
// +k8s:webhook-gen:mutating:group="walruscore.seal.io",version="v1",resource="connectors",scope="Namespaced"
// +k8s:webhook-gen:mutating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="NoneOnDryRun",matchPolicy="Equivalent",timeoutSeconds=10
type ConnectorWebhook struct {
	Client ctrlcli.Client
}

func (r *ConnectorWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.Client = opts.Manager.GetClient()

	return &walruscore.Connector{}, nil
}

var _ ctrlwebhook.CustomValidator = (*ConnectorWebhook)(nil)

func (r *ConnectorWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	conn := obj.(*walruscore.Connector)
	err := func() error {
		// Validate connector name.
		if stringx.StringWidth(conn.Name) > 30 {
			return field.TooLongMaxLength(
				field.NewPath("name"), stringx.StringWidth(conn.Name), 30)
		}
		// Validate connector secret name.
		if conn.Spec.SecretName == "" {
			return field.Required(field.NewPath("spec.secretName"), "transferred secret name is required")
		}
		return nil
	}()
	if err != nil {
		return r.revokeTransferredSecret(ctx, conn), err
	}
	return nil, nil
}

func (r *ConnectorWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	oldConn, newConn := oldObj.(*walruscore.Connector), newObj.(*walruscore.Connector)

	// Validate immutable fields.
	if err := r.validateImmutables(oldConn, newConn); err != nil {
		if oldConn.Spec.SecretName != newConn.Spec.SecretName {
			return r.revokeTransferredSecret(ctx, newConn), err
		}
		return nil, err
	}

	// Validate conditions.
	if err := r.validateConditions(oldConn, newConn); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *ConnectorWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	conn := obj.(*walruscore.Connector)

	// Cannot delete the connector if it is in use.
	cbList := new(walruscore.ConnectorBindingList)
	err := r.Client.List(ctx, cbList,
		ctrlcli.MatchingLabels{
			"spec.connector": kubemeta.GetNamespacedNameKey(conn),
		})
	if err != nil {
		return nil, kerrors.NewInternalError(fmt.Errorf("list related connector binding: %w", err))
	}
	if len(cbList.Items) > 0 {
		return nil, field.Invalid(
			field.NewPath("metadata.name"), conn.GetName(), "blocked by in-used connector bindings")
	}

	return nil, nil
}

func (r *ConnectorWebhook) validateImmutables(oldConn, newConn *walruscore.Connector) error {
	if oldConn.Spec.ApplicableEnvironmentType != newConn.Spec.ApplicableEnvironmentType {
		return field.Invalid(
			field.NewPath("spec.applicableEnvironmentType"), oldConn.Spec.ApplicableEnvironmentType, "field is immutable")
	}
	if oldConn.Spec.Category != newConn.Spec.Category {
		return field.Invalid(
			field.NewPath("spec.category"), oldConn.Spec.Category, "field is immutable")
	}
	if oldConn.Spec.Type != newConn.Spec.Type {
		return field.Invalid(
			field.NewPath("spec.type"), oldConn.Spec.Type, "field is immutable")
	}
	if oldConn.Spec.SecretName != newConn.Spec.SecretName {
		return field.Invalid(
			field.NewPath("spec.secretName"), oldConn.Spec.SecretName, "field is immutable")
	}
	return nil
}

func (r *ConnectorWebhook) validateConditions(oldConn, newConn *walruscore.Connector) error {
	switch {
	case oldConn.Generation != newConn.Generation:
		return nil
	case apistatus.ConnectorConditionConnected.IsUnknown(newConn):
		if len(oldConn.Status.Conditions) == 0 ||
			apistatus.ConnectorConditionConnected.IsTrueOrFalse(oldConn) {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newConn.Status.Conditions, "previous connecting has not been completed yet")
	case apistatus.ConnectorConditionDeleting.Exists(newConn):
		if newConn.DeletionTimestamp != nil {
			return nil
		}
		return field.Invalid(
			field.NewPath("status.conditions"), newConn.Status.Conditions, "deletion has not been requested yet")
	}
	return nil
}

func (r *ConnectorWebhook) Default(ctx context.Context, obj runtime.Object) error {
	conn := obj.(*walruscore.Connector)
	if conn.DeletionTimestamp != nil {
		// Nothing to do for deletion.
		return nil
	}

	// Set applicable environment type.
	if conn.Spec.ApplicableEnvironmentType == "" {
		conn.Spec.ApplicableEnvironmentType = walruscore.EnvironmentTypeDevelopment
	}

	// Transfer configuration to a secret if needed.
	return r.transferSecret(ctx, conn)
}

func (r *ConnectorWebhook) transferSecret(ctx context.Context, conn *walruscore.Connector) error {
	var eSec *core.Secret
	{
		switch {
		case conn.Generation == 0:
			// Create the secret.
			eSec = &core.Secret{
				ObjectMeta: meta.ObjectMeta{
					Namespace:    conn.Namespace,
					GenerateName: "connector-config-",
				},
			}
		case conn.Spec.SecretName != "":
			// Check if we need to update the secret.
			var need bool
			for k := range conn.Spec.Config.Data {
				if !conn.Spec.Config.Data[k].Sensitive {
					continue
				}
				if conn.Spec.Config.Data[k].Value == "(sensitive)" {
					continue
				}
				need = true
				break
			}
			if !need {
				// Return directly if no need to update the secret.
				return nil
			}
			// Update the secret.
			eSec = &core.Secret{
				ObjectMeta: meta.ObjectMeta{
					Namespace: conn.Namespace,
					Name:      conn.Spec.SecretName,
				},
			}
		}
		eSec.Data = make(map[string][]byte, len(conn.Spec.Config.Data))
		for k := range conn.Spec.Config.Data {
			eSec.Data[k] = []byte(conn.Spec.Config.Data[k].Value)
		}
	}
	alignFn := func(aSec *core.Secret) (*core.Secret, bool, error) {
		if aSec.Data == nil {
			aSec.Data = make(map[string][]byte)
		}
		if reflect.DeepEqual(aSec.Data, eSec.Data) {
			return nil, true, nil
		}
		aSec.Data = maps.Clone(eSec.Data)
		return aSec, false, nil
	}
	eSec, err := kubeclientset.CreateWithCtrlClient(ctx, r.Client, eSec,
		kubeclientset.WithUpdateIfExisted(alignFn))
	if err != nil {
		return kerrors.NewInternalError(fmt.Errorf("create connector transferred secret: %w", err))
	}
	conn.Spec.SecretName = eSec.Name

	// Reverse controlling,
	// which ensures the connector not leaked after the secret is deleted.
	kubemeta.ControlOn(conn, eSec, core.SchemeGroupVersion.WithKind("Secret"))

	// Desensitize configuration.
	for k := range conn.Spec.Config.Data {
		if !conn.Spec.Config.Data[k].Sensitive {
			continue
		}
		conn.Spec.Config.Data[k] = walruscore.ConnectorConfigEntry{
			Sensitive: true,
			Value:     "(sensitive)",
		}
	}

	// Overwrite last applied annotation if needed.
	kubemeta.OverwriteLastAppliedAnnotation(conn)

	return nil
}

func (r *ConnectorWebhook) revokeTransferredSecret(ctx context.Context, conn *walruscore.Connector) ctrladmission.Warnings {
	logger := ctrllog.FromContext(ctx)

	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: conn.Namespace,
			Name:      conn.Spec.SecretName,
		},
	}

	err := kubeclientset.DeleteWithCtrlClient(ctx, r.Client, sec,
		kubeclientset.WithDeleteMetaOptions(meta.DeleteOptions{PropagationPolicy: ptr.To(meta.DeletePropagationForeground)}))
	if err != nil {
		logger.Error(err, "clean up connector transferred secret")
		return []string{
			fmt.Sprintf("Unable to revoke the connector transferred secret %q, please remove it manually",
				kubemeta.GetNamespacedNameKey(sec)),
		}
	}

	return nil
}
