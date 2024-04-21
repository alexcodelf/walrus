package walruscore

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlctrl "sigs.k8s.io/controller-runtime/pkg/controller"
	ctrlevent "sigs.k8s.io/controller-runtime/pkg/event"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemmeta"
	tmplfetcher "github.com/seal-io/walrus/pkg/templates/fetcher"
)

// TemplateReconciler reconciles a v1.Template object.
type TemplateReconciler struct {
	Client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*TemplateReconciler)(nil)

// Reconcile reconciles the template.
func (r *TemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	tmpl := new(walruscore.Template)
	err := r.Client.Get(ctx, req.NamespacedName, tmpl)
	if err != nil {
		logger.Error(err, "fetch template")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Clean up if deleted.
	if tmpl.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(tmpl) {
			return ctrl.Result{}, nil
		}

		// List related resource components.
		resCompsList := new(walruscore.ResourceComponentsList)
		err = r.Client.List(ctx, resCompsList,
			ctrlcli.MatchingFields{
				"status.template": kubemeta.GetNamespacedNameKey(tmpl),
			})
		if err != nil {
			logger.Error(err, "list related resource components")
			return ctrl.Result{}, err
		}

		// Notify the template deletion blocking by which resource components.
		if len(resCompsList.Items) > 0 {
			// Notify.
			var msgSb strings.Builder
			msgSb.WriteString("Blocked by resource components: ")
			for i := range resCompsList.Items {
				if i > 0 {
					msgSb.WriteString(", ")
				}
				if i >= 5 {
					msgSb.WriteString("...")
					msgSb.WriteString(fmt.Sprintf(" (total %d)", len(resCompsList.Items)))
					break
				}
				msgSb.WriteString(kubemeta.GetNamespacedNameKey(&resCompsList.Items[i]))
			}
			if !apistatus.TemplateConditionDeleting.IsTrue(tmpl) {
				apistatus.TemplateConditionDeleting.ResetTrue(tmpl, "", msgSb.String())
			} else {
				apistatus.TemplateConditionDeleting.Message(tmpl, msgSb.String())
			}
			tmpl.Status.ConditionSummary = *apistatus.WalkTemplate(&tmpl.Status.StatusDescriptor)

			err = r.Client.Status().Update(ctx, tmpl)
			if err != nil {
				return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
			}

			// Requeue to check if the blocking resource components are deleted or not.
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, tmpl)
		if err != nil {
			logger.Error(err, "unlock template")
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(tmpl) {
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, tmpl)
		if err != nil {
			logger.Error(err, "lock template")
		}
		// NB(thxCode): during concurrent reconciles,
		// release the going handle to avoid a conflict error.
		return ctrl.Result{}, err
	}

	// Skip if already synced.
	if apistatus.TemplateConditionSynced.IsTrueOrFalse(tmpl) {
		return ctrl.Result{}, nil
	}

	// Notify the template is going to sync.
	if !apistatus.TemplateConditionSynced.Exists(tmpl) {
		apistatus.TemplateConditionSynced.Reset(tmpl, "", "Syncing from VCS.")
		tmpl.Status.ConditionSummary = *apistatus.WalkTemplate(&tmpl.Status.StatusDescriptor)
		tmpl.Status.Project = systemmeta.GetProjectName(tmpl.Namespace)

		err = r.Client.Status().Update(ctx, tmpl)
		if err != nil {
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}
	}

	// Fetch template and update status.
	tmpl, err = tmplfetcher.Fetch(ctx, tmpl)
	if err != nil {
		// Notify the template is sync failed.
		logger.Error(err, "fetch template's metadata")
		apistatus.TemplateConditionSynced.False(tmpl, apistatus.TemplateConditionSyncedReasonFailed, err.Error())
	} else {
		// Notify the template is synced.
		apistatus.TemplateConditionSynced.True(tmpl, "", "Synced successfully.")
		tmpl.Status.LastSuccessfulSyncTime = ptr.To(meta.Now())
	}
	tmpl.Status.ConditionSummary = *apistatus.WalkTemplate(&tmpl.Status.StatusDescriptor)
	tmpl.Status.Project = systemmeta.GetProjectName(tmpl.Namespace)

	err = r.Client.Status().Update(ctx, tmpl)
	return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
}

// SetupController sets up the controller.
func (r *TemplateReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &walruscore.ResourceComponents{}, "status.template",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			resComps := obj.(*walruscore.ResourceComponents)
			return []string{resComps.Status.Template.ToNamespacedName().String()}
		})
	if err != nil {
		return fmt.Errorf("index resource components 'status.template': %w", err)
	}

	r.Client = opts.Manager.GetClient()

	// Filter out specific update events of templates.
	tmplFilter := ctrlpredicate.Funcs{
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			oldTmpl, newTmpl := e.ObjectOld.(*walruscore.Template), e.ObjectNew.(*walruscore.Template)
			switch {
			default:
				return false
			case newTmpl.DeletionTimestamp != nil:
				if !systemmeta.IsLocked(newTmpl) {
					return false
				}
				// Deleting during locking.
			case !apistatus.TemplateConditionSynced.Exists(newTmpl):
				// Status is not initialized.
			case apistatus.TemplateConditionSynced.IsTrueOrFalse(oldTmpl) &&
				apistatus.TemplateConditionSynced.IsUnknown(newTmpl):
				// Resync.
			}
			return true
		},
		DeleteFunc: func(e ctrlevent.DeleteEvent) bool {
			// NB(thxCode): during concurrent reconciles,
			// ignore final deletion to avoid a not found error.
			return false
		},
	}

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(
			// Focus on the template.
			&walruscore.Template{},
			ctrlbuilder.WithPredicates(tmplFilter),
		).
		WithOptions(ctrlctrl.Options{
			// Start the same number of reconciles as the number of CPUs,
			// which is able to handle the templates sync from one catalog in parallel.
			// TODO(thxCode): pool this?
			MaxConcurrentReconciles: runtime.GOMAXPROCS(0),
		}).
		Complete(r)
}
