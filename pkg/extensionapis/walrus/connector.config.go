package walrus

import (
	"context"

	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
)

// ConnectorConfigHandler handles v1.ConnectorConfig objects,
// which is a subresource of v1.Connector objects.
type ConnectorConfigHandler struct {
	extensionapi.ObjectInfo
	extensionapi.GetOperation

	Client ctrlcli.Client
}

func newConnectorConfigHandler(opts extensionapi.SetupOptions) *ConnectorConfigHandler {
	h := &ConnectorConfigHandler{}

	// As storage.
	h.ObjectInfo = &walrus.ConnectorConfig{}
	h.GetOperation = extensionapi.WithGet(h)

	// Set client.
	h.Client = opts.Manager.GetClient()

	return h
}

var (
	_ rest.Storage = (*ConnectorConfigHandler)(nil)
	_ rest.Getter  = (*ConnectorConfigHandler)(nil)
)

func (h *ConnectorConfigHandler) New() runtime.Object {
	return &walrus.ConnectorConfig{}
}

func (h *ConnectorConfigHandler) Destroy() {}

func (h *ConnectorConfigHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Get connector.
	conn := new(walruscore.Connector)
	err := h.Client.Get(ctx, key, conn, &opts)
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}

	// Get secret.
	sec := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: key.Namespace,
			Name:      conn.Spec.SecretName,
		},
	}
	err = h.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(sec), sec)
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}

	// Construct connector config.
	connCfg := &walrus.ConnectorConfig{
		ObjectMeta: meta.ObjectMeta{
			Namespace: key.Namespace,
			Name:      key.Name,
		},
		Status: walrus.ConnectorConfigStatus{
			ApplicableEnvironmentType: conn.Spec.ApplicableEnvironmentType,
			Category:                  conn.Spec.Category,
			Type:                      conn.Spec.Type,
			Version:                   conn.Spec.Config.Version,
			Data:                      sec.Data,
			ConditionSummary:          conn.Status.ConditionSummary,
		},
	}
	return connCfg, nil
}
