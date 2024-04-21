package walruscore

import (
	"context"
	"fmt"
	"reflect"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/webhook"
)

// ConnectorBindingWebhook hooks a v1.ConnectorBinding object.
//
// nolint: lll
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="connectorbindings",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
// +k8s:webhook-gen:mutating:group="walruscore.seal.io",version="v1",resource="connectorbindings",scope="Namespaced"
// +k8s:webhook-gen:mutating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="NoneOnDryRun",matchPolicy="Equivalent",timeoutSeconds=10
type ConnectorBindingWebhook struct {
	webhook.DefaultCustomValidator

	Client ctrlcli.Client
}

func (r *ConnectorBindingWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.Client = opts.Manager.GetClient()

	return &walruscore.ConnectorBinding{}, nil
}

var _ ctrlwebhook.CustomValidator = (*ConnectorBindingWebhook)(nil)

func (r *ConnectorBindingWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	cb := obj.(*walruscore.ConnectorBinding)

	conn := &walruscore.Connector{
		ObjectMeta: meta.ObjectMeta{
			Name:      cb.Spec.Connector.Name,
			Namespace: cb.Spec.Connector.Namespace,
		},
	}
	err := r.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(conn), conn)
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("spec.connector"), cb.Spec.Connector, "connector not found")
	}

	// Validate if the given category and type are valid.
	{
		if conn.Spec.Category != cb.Spec.Connector.Category {
			return nil, field.Invalid(
				field.NewPath("spec.connector"), cb.Spec.Connector,
				fmt.Sprintf("connector category mismatch, expected %s, got %s", conn.Spec.Category, cb.Spec.Connector.Category))
		}
		if conn.Spec.Type != cb.Spec.Connector.Type {
			return nil, field.Invalid(
				field.NewPath("spec.connector"), cb.Spec.Connector,
				fmt.Sprintf("connector type mismatch, expected %s, got %s", conn.Spec.Type, cb.Spec.Connector.Type))
		}
	}

	// Validate if the connector be applicable to the type of environment.
	{
		envList := new(walrus.EnvironmentList)
		err = r.Client.List(ctx, envList,
			// Since the environment name is unique,
			// we can use it as the field selector without the namespace(project) of environment.
			ctrlcli.MatchingFields{
				"metadata.name": cb.Namespace,
			})
		if err != nil {
			return nil, err
		}

		if len(envList.Items) != 1 {
			return nil, field.Forbidden(
				field.NewPath("metadata.namespace"), "connector binding must be created in a valid environment namespace")
		}

		if eType, aType := envList.Items[0].Spec.Type, conn.Spec.ApplicableEnvironmentType; eType != aType {
			return nil, field.Invalid(
				field.NewPath("spec.connector"), cb.Spec.Connector,
				fmt.Sprintf("connector is not applicable to this environment, expected %s, got %s", eType, aType))
		}
	}

	// Validate if there is duplicated binding.
	{
		cbList := new(walruscore.ConnectorBindingList)
		err = r.Client.List(ctx, cbList,
			ctrlcli.InNamespace(cb.Namespace))
		if err != nil {
			return nil, err
		}

		for i := range cbList.Items {
			if cbList.Items[i].Spec.Connector.Type == conn.Spec.Type {
				return nil, field.Invalid(
					field.NewPath("spec.connector"), cb.Spec.Connector,
					"connectors for the same purpose cannot be bound repeatedly")
			}
		}
	}

	return nil, nil
}

func (r *ConnectorBindingWebhook) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	oldCb, newCb := oldObj.(*walruscore.ConnectorBinding), newObj.(*walruscore.ConnectorBinding)

	// Validate immutable fields.
	if !reflect.DeepEqual(oldCb.Spec, newCb.Spec) {
		return nil, field.Invalid(
			field.NewPath("spec.connector"), oldCb.Spec.Connector, "field is immutable")
	}

	return nil, nil
}

func (r *ConnectorBindingWebhook) Default(ctx context.Context, obj runtime.Object) error {
	cb := obj.(*walruscore.ConnectorBinding)
	if cb.DeletionTimestamp != nil {
		// Nothing to do for deletion.
		return nil
	}

	// Set the category and type of the connector if needed.
	if cb.Spec.Connector.Category != "" && cb.Spec.Connector.Type != "" {
		return nil
	}
	conn := &walruscore.Connector{
		ObjectMeta: meta.ObjectMeta{
			Name:      cb.Spec.Connector.Name,
			Namespace: cb.Spec.Connector.Namespace,
		},
	}
	err := r.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(conn), conn)
	if err != nil {
		return field.Invalid(
			field.NewPath("spec.connector"), cb.Spec.Connector, "connector not found")
	}
	cb.Spec.Connector.Category = conn.Spec.Category
	cb.Spec.Connector.Type = conn.Spec.Type

	return nil
}
