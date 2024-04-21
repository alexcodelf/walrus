package walruscore

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.uber.org/multierr"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
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
	"github.com/seal-io/walrus/pkg/systemmeta"
	tmpllister "github.com/seal-io/walrus/pkg/templates/lister"
)

// CatalogReconciler reconciles a v1.Catalog object.
type CatalogReconciler struct {
	Client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*CatalogReconciler)(nil)

// Reconcile reconciles the catalog.
func (r *CatalogReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	cat := new(walruscore.Catalog)
	err := r.Client.Get(ctx, req.NamespacedName, cat)
	if err != nil {
		logger.Error(err, "fetch catalog")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Clean up if deleted.
	if cat.DeletionTimestamp != nil {
		// Return if already unlocked.
		if systemmeta.Unlock(cat) {
			return ctrl.Result{}, nil
		}

		// List related templates.
		tmplList := new(walruscore.TemplateList)
		err = r.Client.List(ctx, tmplList,
			ctrlcli.InNamespace(cat.Namespace),
			ctrlcli.MatchingFields{
				"metadata.ownerReferences[controller=true].name": cat.Name,
			})
		if err != nil {
			logger.Error(err, "list related templates")
			return ctrl.Result{}, err
		}

		// Notify the catalog is going to delete,
		// and delete related templates.
		if len(tmplList.Items) > 0 {
			// Notify.
			var msgSb strings.Builder
			msgSb.WriteString("Deleting templates: ")
			for i := range tmplList.Items {
				if i > 0 {
					msgSb.WriteString(", ")
				}
				if i >= 5 {
					msgSb.WriteString("...")
					msgSb.WriteString(fmt.Sprintf(" (total %d)", len(tmplList.Items)))
					break
				}
				msgSb.WriteString(kubemeta.GetNamespacedNameKey(&tmplList.Items[i]))
			}
			if !apistatus.CatalogConditionDeleting.IsTrue(cat) {
				apistatus.CatalogConditionDeleting.ResetTrue(cat, "", msgSb.String())
			} else {
				apistatus.CatalogConditionDeleting.Message(cat, msgSb.String())
			}
			cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
			cat.Status.TemplateCount = int64(len(tmplList.Items))

			err = r.Client.Status().Update(ctx, cat)
			if err != nil {
				return ctrl.Result{}, err
			}

			// Delete.
			rerrs := make([]error, 0, len(tmplList.Items))
			ierrs := make([]error, 0, len(tmplList.Items))
			for i := range tmplList.Items {
				eTmpl := &tmplList.Items[i]
				err = r.Client.Delete(ctx, eTmpl)
				if err != nil && !kerrors.IsNotFound(err) {
					logger.Error(err, "deleted related template", "template", kubemeta.GetNamespacedNameKey(eTmpl))
					if !kerrors.IsInvalid(err) {
						rerrs = append(rerrs, err)
						continue
					}
					ierrs = append(ierrs, err)
				}
			}

			switch {
			case len(rerrs) != 0:
				return ctrl.Result{}, multierr.Combine(rerrs...)
			case len(ierrs) != 0:
				msgSb = strings.Builder{}
				msgSb.WriteString("Blocked by templates: ")
				for i := range ierrs {
					if i > 0 {
						msgSb.WriteString(", ")
					}
					if i >= 5 {
						msgSb.WriteString("...")
						msgSb.WriteString(fmt.Sprintf(" (total %d)", len(ierrs)))
						break
					}
					msgSb.WriteString(ierrs[i].Error())
				}
				apistatus.CatalogConditionDeleting.Message(cat, msgSb.String())
				cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
				err = r.Client.Status().Update(ctx, cat)
				if err != nil {
					return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
				}

				// Requeue to check if the blocking templates are be able to delete or not.
				return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
			}

			return ctrl.Result{}, nil
		}

		// Unlock.
		_, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, cat)
		if err != nil {
			logger.Error(err, "unlock catalog")
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		return ctrl.Result{}, nil
	}

	// Lock if not.
	if !systemmeta.Lock(cat) {
		cat, err = kubeclientset.UpdateWithCtrlClient(ctx, r.Client, cat)
		if err != nil {
			logger.Error(err, "lock catalog")
			return ctrl.Result{}, err
		}
	}

	// Skip if fetch failed.
	if apistatus.CatalogConditionFetched.IsFalse(cat) {
		return ctrl.Result{}, nil
	}

	// Notify the catalog is going to fetch.
	if !apistatus.CatalogConditionFetched.Exists(cat) {
		apistatus.CatalogConditionFetched.Reset(cat, "", "Fetching templates from VCS.")
		cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
		cat.Status.Project = systemmeta.GetProjectName(cat.Namespace)
		cat.Status.URL = cat.Spec.VCSSource.URL

		err = r.Client.Status().Update(ctx, cat)
		if err != nil {
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}
	}

	// Fetch templates if needed.
	if apistatus.CatalogConditionFetched.IsUnknown(cat) {
		// Fetch templates.
		tmpls, err := tmpllister.List(ctx, cat)
		if err != nil {
			logger.Error(err, "fetch templates' metadata")

			// Notify the catalog is fetch failed.
			apistatus.CatalogConditionFetched.False(cat, apistatus.CatalogConditionFetchedReasonAllFailed, err.Error())
			cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
			cat.Status.Project = systemmeta.GetProjectName(cat.Namespace)
			cat.Status.URL = cat.Spec.VCSSource.URL

			err = r.Client.Status().Update(ctx, cat)
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		// Update or create templates.
		errs := make([]error, 0, len(tmpls))
		for i := range tmpls {
			eTmpl := &tmpls[i]
			kubemeta.ControlOn(eTmpl, cat, walruscore.SchemeGroupVersion.WithKind("Catalog"))

			_, err = kubeclientset.CreateWithCtrlClient(ctx, r.Client, eTmpl)
			if err != nil {
				logger.Error(err, "update or create template", "template", kubemeta.GetNamespacedNameKey(eTmpl))
				errs = append(errs, err)
			}
		}
		if len(errs) != 0 {
			err = multierr.Combine(errs...)

			// Notify the catalog is fetch failed.
			rea := apistatus.CatalogConditionFetchedReasonPartialFailed
			if len(tmpls) == len(errs) {
				rea = apistatus.CatalogConditionFetchedReasonAllFailed
			}
			apistatus.CatalogConditionFetched.False(cat, rea, err.Error())
			cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
			cat.Status.Project = systemmeta.GetProjectName(cat.Namespace)
			cat.Status.URL = cat.Spec.VCSSource.URL

			err = r.Client.Status().Update(ctx, cat)
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		// Notify the catalog is fetched.
		apistatus.CatalogConditionFetched.True(cat, "", fmt.Sprintf("Fetched %d templates from VCS.", len(tmpls)))
		cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
		cat.Status.Project = systemmeta.GetProjectName(cat.Namespace)
		cat.Status.URL = cat.Spec.VCSSource.URL

		err = r.Client.Status().Update(ctx, cat)
		if err != nil {
			return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
		}

		// NB(thxCode): Go ahead to watch if templates are synced.
	}

	// List related templates.
	tmplList := new(walruscore.TemplateList)
	err = r.Client.List(ctx, tmplList,
		ctrlcli.InNamespace(cat.Namespace),
		ctrlcli.MatchingFields{
			"metadata.ownerReferences[controller=true].name": cat.Name,
		})
	if err != nil {
		logger.Error(err, "list related templates")
		return ctrl.Result{}, err
	}

	// Count related templates' status.
	var done, failed, remain int
	for i := range tmplList.Items {
		eTmpl := &tmplList.Items[i]
		switch walruscore.ConditionStatus(apistatus.TemplateConditionSynced.GetStatus(eTmpl)) {
		case walruscore.ConditionTrue:
			done++
		case walruscore.ConditionFalse:
			failed++
		}
	}
	remain = len(tmplList.Items) - done - failed

	// Notify the catalog's syncing result.
	switch {
	case remain != 0:
		apistatus.CatalogConditionSyncedTemplates.Unknown(cat,
			apistatus.CatalogConditionSyncedTemplatesReasonSyncing,
			fmt.Sprintf("Synced %d templates, remaining %d templates.", done+failed, remain))
	case failed != 0:
		apistatus.CatalogConditionSyncedTemplates.False(cat,
			apistatus.CatalogConditionSyncedTemplatesReasonFailed,
			fmt.Sprintf("Synced all templates, but failed %d templates.", failed))
	default:
		apistatus.CatalogConditionSyncedTemplates.True(cat, "",
			fmt.Sprintf("Synced %d templates successfully.", done))
		cat.Status.LastSuccessfulSyncTime = ptr.To(meta.Now())
	}
	cat.Status.ConditionSummary = *apistatus.WalkCatalog(&cat.Status.StatusDescriptor)
	cat.Status.Project = systemmeta.GetProjectName(cat.Namespace)
	cat.Status.TemplateCount = int64(len(tmplList.Items))

	err = r.Client.Status().Update(ctx, cat)
	return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
}

// SetupController sets up the controller.
func (r *CatalogReconciler) SetupController(ctx context.Context, opts controller.SetupOptions) error {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err := fi.IndexField(ctx, &walruscore.Template{}, "metadata.ownerReferences[controller=true].name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			ctrler := kubemeta.GetControllerOfNoCopy(obj)
			if ctrler == nil {
				return nil
			}
			if ctrler.APIVersion != walruscore.SchemeGroupVersion.String() ||
				ctrler.Kind != "Catalog" {
				return nil
			}
			return []string{ctrler.Name}
		})
	if err != nil {
		return fmt.Errorf("index template 'metadata.ownerReferences[controller=true].name': %w", err)
	}

	r.Client = opts.Manager.GetClient()

	// Filter out specific update events of catalogs.
	catFilter := ctrlpredicate.Funcs{
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			oldCat, newCat := e.ObjectOld.(*walruscore.Catalog), e.ObjectNew.(*walruscore.Catalog)
			switch {
			default:
				return false
			case newCat.DeletionTimestamp != nil:
				if !systemmeta.IsLocked(newCat) {
					return false
				}
				// Deleting during locking.
			case newCat.Generation != oldCat.Generation && !reflect.DeepEqual(newCat.Spec.Filters, oldCat.Spec.Filters):
				// Interested mutable fields is changed.
			case !apistatus.CatalogConditionFetched.Exists(newCat):
				// Status is not initialized.
			case !apistatus.CatalogConditionSyncedTemplates.Exists(newCat):
				// Status is not completed.
			}
			return true
		},
	}

	// Filter out deleting or syncing update events of templates.
	tmplFilter := ctrlpredicate.Funcs{
		CreateFunc: func(_ ctrlevent.CreateEvent) bool {
			// Ignore creation.
			return false
		},
		GenericFunc: func(_ ctrlevent.GenericEvent) bool {
			// Ignore generation.
			return false
		},
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
			case apistatus.TemplateConditionSynced.IsUnknown(oldTmpl) &&
				apistatus.TemplateConditionSynced.IsTrueOrFalse(newTmpl):
				// Synced.
			}
			return true
		},
	}

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(
			// Focus on the catalog.
			&walruscore.Catalog{},
			ctrlbuilder.WithPredicates(catFilter),
		).
		Owns(
			// Own the templates.
			&walruscore.Template{},
			ctrlbuilder.WithPredicates(tmplFilter),
		).
		Complete(r)
}
