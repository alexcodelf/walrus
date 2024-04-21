package walruscore

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlevent "sigs.k8s.io/controller-runtime/pkg/event"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ConnectorReconciler reconciles a v1.Connector object.
type ConnectorReconciler struct {
	Client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ConnectorReconciler)(nil)

func (r *ConnectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	conn := new(walruscore.Connector)
	err := r.Client.Get(ctx, req.NamespacedName, conn)
	if err != nil {
		logger.Error(err, "fetch connector")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Clean up if deleted.
	if conn.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(conn) {
			return ctrl.Result{}, nil
		}

		// List related connector bindings.
		cbList := new(walruscore.ConnectorBindingList)
		err = r.Client.List(ctx, cbList,
			ctrlcli.MatchingFields{
				"spec.connector": kubemeta.GetNamespacedNameKey(conn),
			})
		if err != nil {
			logger.Error(err, "list related connector binding")
			return ctrl.Result{}, err
		}

		// Notify the connector deletion blocking by which connector binding.
		if len(cbList.Items) > 0 {
			// Notify.
			var msgSb strings.Builder
			msgSb.WriteString("Blocked by connector binding: ")
			for i := range cbList.Items {
				if i > 0 {
					msgSb.WriteString(", ")
				}
				if i >= 5 {
					msgSb.WriteString("...")
					msgSb.WriteString(fmt.Sprintf(" (total %d)", len(cbList.Items)))
					break
				}
				msgSb.WriteString(kubemeta.GetNamespacedNameKey(&cbList.Items[i]))
			}
			if !apistatus.ConnectorConditionDeleting.IsTrue(conn) {
				apistatus.ConnectorConditionDeleting.ResetTrue(conn, "", msgSb.String())
			} else {
				apistatus.ConnectorConditionDeleting.Message(conn, msgSb.String())
			}
			conn.Status.ConditionSummary = *apistatus.WalkConnector(&conn.Status.StatusDescriptor)

			err = r.Client.Status().Update(ctx, conn)
			if err != nil {
				return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
			}

			// Requeue to check if the blocking connector bindings are deleted or not.
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}

		// Delete related secret.
		sec := &core.Secret{
			ObjectMeta: meta.ObjectMeta{
				Namespace: conn.Namespace,
				Name:      conn.Spec.SecretName,
			},
		}
		err = kubeclientset.DeleteWithCtrlClient(ctx, r.Client, sec)
		if err != nil {
			logger.Error(err, "delete related secret")
			return ctrl.Result{}, err
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, conn)
		if err != nil {
			logger.Error(err, "unlock connector")
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(conn) {
		conn, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, conn)
		if err != nil {
			logger.Error(err, "lock connector")
			return ctrl.Result{}, err
		}
	}

	// Skip if already synced.
	if apistatus.ConnectorConditionConnected.IsTrueOrFalse(conn) {
		return ctrl.Result{}, nil
	}

	// Notify the connector is going to dail.
	if !apistatus.ConnectorConditionConnected.Exists(conn) {
		apistatus.ConnectorConditionConnected.Reset(conn, "", "Dialing.")
		conn.Status.ConditionSummary = *apistatus.WalkConnector(&conn.Status.StatusDescriptor)
		conn.Status.Project = systemmeta.GetProjectName(conn.Namespace)

		err = r.Client.Status().Update(ctx, conn)
		if err != nil {
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}
	}

	// Get resource handler.
	rh, err := resourcehandlers.Get(ctx, resourcehandler.GetOptions{
		Connector: walruscore.ConnectorReference{
			Namespace: conn.Namespace,
			Name:      conn.Name,
		},
		Type: conn.Spec.Type,
	})
	if err != nil {
		logger.Error(err, "get resource handler")
		return ctrl.Result{}, err
	}

	// Dial with resource handler and update status.
	if err = rh.IsConnected(ctx); err != nil {
		// Notify the connector is disconnected.
		logger.Error(err, "disconnected")
		apistatus.ConnectorConditionConnected.False(conn, apistatus.ConnectorConditionConnectedReasonFailed, err.Error())
	} else {
		// Notify the connector is connected.
		apistatus.ConnectorConditionConnected.True(conn, "", "Dialed successfully.")
	}
	conn.Status.ConditionSummary = *apistatus.WalkConnector(&conn.Status.StatusDescriptor)
	conn.Status.Project = systemmeta.GetProjectName(conn.Namespace)

	err = r.Client.Status().Update(ctx, conn)
	return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
}

func (r *ConnectorReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &walruscore.ConnectorBinding{}, "spec.connector",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			cb := obj.(*walruscore.ConnectorBinding)
			return []string{cb.Spec.Connector.ToNamespacedName().String()}
		})
	if err != nil {
		return fmt.Errorf("index connector binding 'spec.connector': %w", err)
	}

	r.Client = opts.Manager.GetClient()

	// Filter out specific update events of connectors.
	connFilter := ctrlpredicate.Funcs{
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			oldConn, newConn := e.ObjectOld.(*walruscore.Connector), e.ObjectNew.(*walruscore.Connector)
			switch {
			default:
				return false
			case newConn.DeletionTimestamp != nil:
				if !systemmeta.IsLocked(newConn) {
					return false
				}
				// Deleting during locking.
			case newConn.Generation != oldConn.Generation && !reflect.DeepEqual(newConn.Spec.Config, oldConn.Spec.Config):
				// Interested mutable fields is changed.
			case !apistatus.ConnectorConditionConnected.Exists(newConn):
				// Status is not initialized.
			case apistatus.ConnectorConditionConnected.IsTrueOrFalse(oldConn) &&
				apistatus.ConnectorConditionConnected.IsUnknown(newConn):
				// Reconnect.
			}
			return true
		},
	}

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(
			// Focus on the connector.
			&walruscore.Connector{},
			ctrlbuilder.WithPredicates(connFilter),
		).
		Complete(r)
}
