package walruscore

import (
	"context"
	"fmt"
	"maps"
	"reflect"
	"strings"
	"time"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlevent "sigs.k8s.io/controller-runtime/pkg/event"
	ctrlhandler "sigs.k8s.io/controller-runtime/pkg/handler"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// ConnectorBindingReconciler reconciles a v1.ConnectorBinding object.
type ConnectorBindingReconciler struct {
	Client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ConnectorBindingReconciler)(nil)

func (r *ConnectorBindingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	cb := new(walruscore.ConnectorBinding)
	err := r.Client.Get(ctx, req.NamespacedName, cb)
	if err != nil {
		logger.Error(err, "fetch connector binding")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Clean up if deleted.
	if cb.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(cb) {
			return ctrl.Result{}, nil
		}

		// Get related environment.
		envList := &walrus.EnvironmentList{}
		err = r.Client.List(ctx, envList,
			// Since the environment name is unique,
			// we can use it as the field selector without the namespace(project) of environment.
			ctrlcli.MatchingFields{
				"metadata.name": cb.Namespace,
			})
		if err != nil {
			logger.Error(err, "get related environment")
			return ctrl.Result{}, err
		}

		// Unlabel environment if needed.
		if len(envList.Items) == 1 && envList.Items[0].DeletionTimestamp == nil {
			env := &envList.Items[0]
			lbs := maps.Clone(env.Labels)
			delete(lbs, getConnectorBoundTypeLabel(cb))
			if len(lbs) != len(env.Labels) {
				env.Labels = lbs
				err = r.Client.Update(ctx, env)
				if err != nil && !kerrors.IsNotFound(err) {
					logger.Error(err, "unlabel environment", "environment", kubemeta.GetNamespacedNameKey(env))
					return ctrl.Result{}, err
				}
			}
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, cb)
		if err != nil {
			logger.Error(err, "unlock connector binding")
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(cb) {
		cb, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, cb)
		if err != nil {
			logger.Error(err, "lock connector binding")
			return ctrl.Result{}, err
		}
	}

	// Get related environment.
	envList := &walrus.EnvironmentList{}
	err = r.Client.List(ctx, envList,
		// Since the environment name is unique,
		// we can use it as the field selector without the namespace(project) of environment.
		ctrlcli.MatchingFields{
			"metadata.name": cb.Namespace,
		})
	if err != nil {
		logger.Error(err, "get related environment")
		return ctrl.Result{}, err
	}

	if len(envList.Items) != 1 {
		// NB(thxCode): we should never reach here.
		logger.Error(nil, "cannot fetch related environment")
		return ctrl.Result{RequeueAfter: time.Second}, nil
	}

	// Label environment if needed.
	env := &envList.Items[0]
	lbs := maps.Clone(env.Labels)
	if lbs == nil {
		lbs = make(map[string]string)
	}
	lbs[getConnectorBoundTypeLabel(cb)] = ""
	if len(lbs) != len(env.Labels) {
		env.Labels = lbs
		err = r.Client.Update(ctx, env)
		if err != nil {
			logger.Error(err, "label environment", "environment", kubemeta.GetNamespacedNameKey(env))
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *ConnectorBindingReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &walrus.Environment{}, "metadata.name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return fmt.Errorf("index environment 'metadata.name': %w", err)
	}

	// Filter out specific update events of connector bindings.
	cbFilter := ctrlpredicate.Funcs{
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			oldCb, newCb := e.ObjectOld.(*walruscore.ConnectorBinding), e.ObjectNew.(*walruscore.ConnectorBinding)
			return !reflect.DeepEqual(oldCb.Spec, newCb.Spec)
		},
	}

	// Filter out updating environment.
	envFilter := ctrlpredicate.Not(ctrlpredicate.Funcs{
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			return false
		},
	})

	r.Client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(
			// Focus on the connector binding.
			&walruscore.ConnectorBinding{},
			ctrlbuilder.WithPredicates(cbFilter),
		).
		Watches(
			// Requeue when updating an Environment.
			&walrus.Environment{},
			ctrlhandler.EnqueueRequestsFromMapFunc(r.findObjectsWhenEnvironmentUpdating),
			ctrlbuilder.WithPredicates(envFilter),
		).
		Complete(r)
}

func (r *ConnectorBindingReconciler) findObjectsWhenEnvironmentUpdating(ctx context.Context, env ctrlcli.Object) []ctrlreconcile.Request {
	logger := ctrllog.FromContext(ctx)

	cbList := new(walruscore.ConnectorBindingList)
	err := r.Client.List(ctx, cbList,
		ctrlcli.InNamespace(env.GetName()))
	if err != nil {
		logger.Error(err, "list connector bindings")
		return nil
	}

	lbs := sets.KeySet(env.GetLabels())

	reqs := make([]ctrlreconcile.Request, 0, len(cbList.Items))
	for i := range cbList.Items {
		if lbs.Has(getConnectorBoundTypeLabel(&cbList.Items[i])) {
			continue
		}
		reqs = append(reqs, ctrlreconcile.Request{
			NamespacedName: ctrlcli.ObjectKey{
				Namespace: cbList.Items[i].Namespace,
				Name:      cbList.Items[i].Name,
			},
		})
	}
	return reqs
}

const ConnectorBoundTypeLabelPrefix = "bound.connector.walrus.seal.io/type-"

func getConnectorBoundTypeLabel(cb *walruscore.ConnectorBinding) string {
	return ConnectorBoundTypeLabelPrefix + strings.ToLower(cb.Spec.Connector.Type)
}
