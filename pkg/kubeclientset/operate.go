package kubeclientset

import (
	"context"
	"errors"
	"reflect"
	"time"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	// MetaObject is the interface for the object with metadata.
	MetaObject = ctrlcli.Object

	// AlignWithFn is a function to compare the actual object with the expected object,
	// and returns aligned object if the actual object is no the same with the expected object.
	AlignWithFn[T MetaObject] func(actualCopied T) (aligned T, skip bool, err error)

	// CompareWithFn is a function to compare the actual object with the expected object,
	// and returns true if the actual object is the same with the expected object.
	CompareWithFn[T MetaObject] func(actualCopied T) (skip bool)
)

type (
	CreateClient[T MetaObject] interface {
		Create(ctx context.Context, obj T, opts meta.CreateOptions) (T, error)
		Update(ctx context.Context, obj T, opts meta.UpdateOptions) (T, error)
		Get(ctx context.Context, name string, opts meta.GetOptions) (T, error)
		Delete(ctx context.Context, name string, opts meta.DeleteOptions) error
	}

	_CreateOptions[T MetaObject] struct {
		meta.CreateOptions
		UpdateAlignFunc     AlignWithFn[T]
		RecreateCompareFunc CompareWithFn[T]
	}

	CreateOption[T MetaObject] func(*_CreateOptions[T])
)

// WithCreateMetaOptions sets the create options.
func WithCreateMetaOptions[T MetaObject](opts meta.CreateOptions) CreateOption[T] {
	return func(co *_CreateOptions[T]) {
		co.CreateOptions = opts
	}
}

// WithUpdateIfExisted with the align function to update the resource if existed.
//
// WithUpdateIfExisted is conflict to WithRecreateIfDuplicated, if both provided,
// WithUpdateIfExisted will be used.
func WithUpdateIfExisted[T MetaObject](fn AlignWithFn[T]) CreateOption[T] {
	return func(co *_CreateOptions[T]) {
		co.UpdateAlignFunc = fn
	}
}

// WithRecreateIfDuplicated with the compare function to recreate the resource if different.
//
// WithRecreateIfDuplicated is conflict to WithUpdateIfExisted, if both provided,
// WithUpdateIfExisted will be used.
func WithRecreateIfDuplicated[T MetaObject](fn CompareWithFn[T]) CreateOption[T] {
	return func(co *_CreateOptions[T]) {
		co.RecreateCompareFunc = fn
	}
}

// Create is similar to Apply, will create the resource if it does not exist.
//
// Create updates the resource if WithUpdateIfExisted provided,
// or recreate the resource if WithRecreateIfDuplicated provided.
// Select one from WithUpdateIfExisted and WithRecreateIfDuplicated, if both provided,
// WithUpdateIfExisted will be used.
func Create[T MetaObject](ctx context.Context, cli CreateClient[T], expected T, opts ...CreateOption[T]) (T, error) {
	if reflect.ValueOf(expected).IsZero() {
		return expected, errors.New("expected is nil")
	}

	var co _CreateOptions[T]
	for i := range opts {
		opts[i](&co)
	}

	var (
		name   = expected.GetName()
		err    = errors.New("resource name may not be empty")
		actual T
	)

	if name != "" {
		actual, err = cli.Get(ctx, name, meta.GetOptions{
			ResourceVersion: "0",
		})
		if err != nil && !kerrors.IsNotFound(err) {
			return actual, err
		}
	}

	// Create if not found or deleting.
	if err != nil || actual.GetDeletionTimestamp() != nil {
		deleting := actual.GetDeletionTimestamp() != nil && len(actual.GetFinalizers()) == 0
		if deleting {
			// NB(thxCode): sleep a while to avoid server flipping.
			time.Sleep(10 * time.Millisecond)
		}
		actual, err = cli.Create(ctx, expected, meta.CreateOptions{
			DryRun: co.DryRun,
		})
		if err != nil && kerrors.IsAlreadyExists(err) {
			// Retry on already existed if:
			// - configure align function.
			// - configure compare function.
			// - the resource is deleting without finalizers.
			if co.UpdateAlignFunc != nil || co.RecreateCompareFunc != nil || deleting {
				return Create(ctx, cli, expected, opts...)
			}
			err = nil
		}
		return actual, err
	}

	switch {
	case co.UpdateAlignFunc != nil:
		var (
			copied T
			skip   bool
		)
		copied, skip, err = co.UpdateAlignFunc(actual.DeepCopyObject().(T))
		if err != nil {
			return actual, err
		}
		if skip {
			return actual, nil
		}

		// Copy resource version for update.
		//
		// And keep the original labels, annotations, finalizers, and owner references if they are not set.
		// If you want to clean the above fields, please set them to empty in the expected object.
		copiedOm, actualOm := copied, actual
		copiedOm.SetResourceVersion(actualOm.GetResourceVersion())
		if copiedOm.GetLabels() == nil {
			copiedOm.SetLabels(actualOm.GetLabels())
		}
		if copiedOm.GetAnnotations() == nil {
			copiedOm.SetAnnotations(actualOm.GetAnnotations())
		}
		if copiedOm.GetFinalizers() == nil {
			copiedOm.SetFinalizers(actualOm.GetFinalizers())
		}
		if copiedOm.GetOwnerReferences() == nil {
			copiedOm.SetOwnerReferences(actualOm.GetOwnerReferences())
		}

		updated, err := cli.Update(ctx, copied, meta.UpdateOptions{
			DryRun: co.DryRun,
		})
		if err == nil || !kerrors.IsConflict(err) || !kerrors.IsNotAcceptable(err) {
			return updated, err
		}

		// Retry if conflicted.
		return Create(ctx, cli, expected, opts...)
	case co.RecreateCompareFunc != nil:
		skip := co.RecreateCompareFunc(actual.DeepCopyObject().(T))
		if skip {
			return actual, nil
		}

		err = cli.Delete(ctx, name, meta.DeleteOptions{
			DryRun:            co.DryRun,
			PropagationPolicy: ptr.To(meta.DeletePropagationForeground),
		})
		if err != nil && !kerrors.IsNotFound(err) {
			return actual, err
		}

		// Recreate.
		return Create(ctx, cli, expected, opts...)
	}

	return actual, nil
}

// CreateWithCtrlClient is similar to Create, but uses the ctrl client.
func CreateWithCtrlClient[T MetaObject](ctx context.Context, cli ctrlcli.Client, expected T, opts ...CreateOption[T]) (T, error) {
	if reflect.ValueOf(expected).IsZero() {
		return expected, errors.New("expected is nil")
	}

	var co _CreateOptions[T]
	for i := range opts {
		opts[i](&co)
	}

	var (
		name   = expected.GetName()
		err    = errors.New("resource name may not be empty")
		actual = expected.DeepCopyObject().(T)
	)

	if name != "" {
		err = cli.Get(ctx, ctrlcli.ObjectKeyFromObject(expected), actual)
		if err != nil && !kerrors.IsNotFound(err) {
			return actual, err
		}
	}

	// Create if not found or deleting.
	if err != nil || actual.GetDeletionTimestamp() != nil {
		deleting := actual.GetDeletionTimestamp() != nil && len(actual.GetFinalizers()) == 0
		if deleting {
			// NB(thxCode): sleep a while to avoid server flipping.
			time.Sleep(10 * time.Millisecond)
		}
		err = cli.Create(ctx, expected, &ctrlcli.CreateOptions{
			DryRun:       co.DryRun,
			FieldManager: co.FieldManager,
			Raw:          ptr.To(co.CreateOptions),
		})
		if err == nil {
			return expected, nil
		}
		if kerrors.IsAlreadyExists(err) {
			// Retry on already existed if:
			// - configure align function.
			// - configure compare function.
			// - the resource is deleting without finalizers.
			if co.UpdateAlignFunc != nil || co.RecreateCompareFunc != nil || deleting {
				return CreateWithCtrlClient(ctx, cli, expected, opts...)
			}
			err = nil
		}
		return actual, err
	}

	switch {
	case co.UpdateAlignFunc != nil:
		var (
			copied T
			skip   bool
		)
		copied, skip, err = co.UpdateAlignFunc(actual.DeepCopyObject().(T))
		if err != nil {
			return actual, err
		}
		if skip {
			return actual, nil
		}

		// Copy resource version for update.
		//
		// And keep the original labels, annotations, finalizers, and owner references if they are not set.
		// If you want to clean the above fields, please set them to empty in the expected object.
		copiedOm, actualOm := copied, actual
		copiedOm.SetResourceVersion(actualOm.GetResourceVersion())
		if copiedOm.GetLabels() == nil {
			copiedOm.SetLabels(actualOm.GetLabels())
		}
		if copiedOm.GetAnnotations() == nil {
			copiedOm.SetAnnotations(actualOm.GetAnnotations())
		}
		if copiedOm.GetFinalizers() == nil {
			copiedOm.SetFinalizers(actualOm.GetFinalizers())
		}
		if copiedOm.GetOwnerReferences() == nil {
			copiedOm.SetOwnerReferences(actualOm.GetOwnerReferences())
		}

		err = cli.Update(ctx, copied, &ctrlcli.UpdateOptions{
			DryRun: co.DryRun,
			Raw:    &meta.UpdateOptions{DryRun: co.DryRun},
		})
		if err == nil || !kerrors.IsConflict(err) || !kerrors.IsNotAcceptable(err) {
			return copied, err
		}

		// Retry if conflicted.
		return CreateWithCtrlClient(ctx, cli, expected, opts...)
	case co.RecreateCompareFunc != nil:
		skip := co.RecreateCompareFunc(actual.DeepCopyObject().(T))
		if skip {
			return actual, nil
		}

		err = cli.Delete(ctx, expected, &ctrlcli.DeleteOptions{
			DryRun:            co.DryRun,
			PropagationPolicy: ptr.To(meta.DeletePropagationForeground),
		})
		if err != nil && !kerrors.IsNotFound(err) {
			return actual, err
		}

		// Recreate.
		return CreateWithCtrlClient(ctx, cli, expected, opts...)
	}

	return actual, nil
}

type (
	UpdateClient[T MetaObject] interface {
		Create(ctx context.Context, obj T, opts meta.CreateOptions) (T, error)
		Update(ctx context.Context, obj T, opts meta.UpdateOptions) (T, error)
		Get(ctx context.Context, name string, opts meta.GetOptions) (T, error)
	}

	_UpdateOptions[T MetaObject] struct {
		meta.UpdateOptions
		AlignFunc          AlignWithFn[T]
		CreateIfNotExisted bool
	}

	UpdateOption[T MetaObject] func(*_UpdateOptions[T])
)

// WithUpdateMetaOptions sets the update options.
func WithUpdateMetaOptions[T MetaObject](opts meta.UpdateOptions) UpdateOption[T] {
	return func(uo *_UpdateOptions[T]) {
		uo.UpdateOptions = opts
	}
}

// WithUpdateAlign with the align function to update the resource.
func WithUpdateAlign[T MetaObject](fn AlignWithFn[T]) UpdateOption[T] {
	return func(uo *_UpdateOptions[T]) {
		uo.AlignFunc = fn
	}
}

// WithCreateIfNotExisted will create the resource if it does not exist.
func WithCreateIfNotExisted[T MetaObject]() UpdateOption[T] {
	return func(uo *_UpdateOptions[T]) {
		uo.CreateIfNotExisted = true
	}
}

// Update will update the resource if it exists,
// and returns the updated resource.
//
// Update returns error if the resource is not found or updating failed.
//
// Update will retry if the resource is updating conflicted when AlignWithFn is provided.
func Update[T MetaObject](ctx context.Context, cli UpdateClient[T], expected T, opts ...UpdateOption[T]) (T, error) {
	if reflect.ValueOf(expected).IsZero() {
		return expected, errors.New("expected is nil")
	}

	var uo _UpdateOptions[T]
	for i := range opts {
		opts[i](&uo)
	}

	name := expected.GetName()
	if name == "" {
		return expected, errors.New("resource name may not be empty")
	}

	actual, err := cli.Get(ctx, name, meta.GetOptions{
		ResourceVersion: "0",
	})
	if err != nil {
		if kerrors.IsNotFound(err) && uo.CreateIfNotExisted {
			actual, err = cli.Create(ctx, expected, meta.CreateOptions{
				DryRun: uo.DryRun,
			})
			if err != nil && kerrors.IsAlreadyExists(err) {
				// Retry if already existed.
				return Update(ctx, cli, expected, opts...)
			}
		}
		return actual, err
	}

	var copied T
	if uo.AlignFunc != nil {
		var skip bool
		copied, skip, err = uo.AlignFunc(actual.DeepCopyObject().(T))
		if err != nil {
			return actual, err
		}
		if skip {
			return actual, nil
		}
	} else {
		copied = expected.DeepCopyObject().(T)
		// Copy resource version for update.
		//
		// And keep the original labels, annotations, finalizers, and owner references if they are not set.
		// If you want to clean the above fields, please set them to empty in the expected object.
		copiedOm, actualOm := copied, actual
		copiedOm.SetResourceVersion(actualOm.GetResourceVersion())
		if copiedOm.GetLabels() == nil {
			copiedOm.SetLabels(actualOm.GetLabels())
		}
		if copiedOm.GetAnnotations() == nil {
			copiedOm.SetAnnotations(actualOm.GetAnnotations())
		}
		if copiedOm.GetFinalizers() == nil {
			copiedOm.SetFinalizers(actualOm.GetFinalizers())
		}
		if copiedOm.GetOwnerReferences() == nil {
			copiedOm.SetOwnerReferences(actualOm.GetOwnerReferences())
		}
	}

	updated, err := cli.Update(ctx, copied, meta.UpdateOptions{
		DryRun: uo.DryRun,
	})
	if err != nil {
		if !kerrors.IsConflict(err) && !kerrors.IsNotAcceptable(err) {
			return actual, err
		}

		// Retry if conflicted when align function is provided.
		if uo.AlignFunc != nil {
			return Update(ctx, cli, expected, opts...)
		}
	}

	return updated, err
}

// UpdateWithCtrlClient is similar to Update, but uses the ctrl client.
func UpdateWithCtrlClient[T MetaObject](ctx context.Context, cli ctrlcli.Client, expected T, opts ...UpdateOption[T]) (T, error) {
	if reflect.ValueOf(expected).IsZero() {
		return expected, errors.New("expected is nil")
	}

	var uo _UpdateOptions[T]
	for i := range opts {
		opts[i](&uo)
	}

	name := expected.GetName()
	if name == "" {
		return expected, errors.New("resource name may not be empty")
	}

	actual := expected.DeepCopyObject().(T)
	err := cli.Get(ctx, ctrlcli.ObjectKeyFromObject(expected), actual)
	if err != nil {
		if kerrors.IsNotFound(err) && uo.CreateIfNotExisted {
			actual = expected.DeepCopyObject().(T)
			err = cli.Create(ctx, actual, &ctrlcli.CreateOptions{
				DryRun:       uo.DryRun,
				FieldManager: uo.FieldManager,
			})
			if err != nil && kerrors.IsAlreadyExists(err) {
				// Retry if already existed.
				return UpdateWithCtrlClient(ctx, cli, expected, opts...)
			}
		}
		return actual, err
	}

	var copied T
	if uo.AlignFunc != nil {
		var skip bool
		copied, skip, err = uo.AlignFunc(actual.DeepCopyObject().(T))
		if err != nil {
			return actual, err
		}
		if skip {
			return actual, nil
		}
	} else {
		copied = expected.DeepCopyObject().(T)
		// Copy resource version for update.
		//
		// And keep the original labels, annotations, finalizers, and owner references if they are not set.
		// If you want to clean the above fields, please set them to empty in the expected object.
		copiedOm, actualOm := copied, actual
		copiedOm.SetResourceVersion(actualOm.GetResourceVersion())
		if copiedOm.GetLabels() == nil {
			copiedOm.SetLabels(actualOm.GetLabels())
		}
		if copiedOm.GetAnnotations() == nil {
			copiedOm.SetAnnotations(actualOm.GetAnnotations())
		}
		if copiedOm.GetFinalizers() == nil {
			copiedOm.SetFinalizers(actualOm.GetFinalizers())
		}
		if copiedOm.GetOwnerReferences() == nil {
			copiedOm.SetOwnerReferences(actualOm.GetOwnerReferences())
		}
	}

	updated := copied
	err = cli.Update(ctx, updated, &ctrlcli.UpdateOptions{
		DryRun:       uo.DryRun,
		FieldManager: uo.FieldManager,
		Raw:          ptr.To(uo.UpdateOptions),
	})
	if err != nil {
		if !kerrors.IsConflict(err) && !kerrors.IsNotAcceptable(err) {
			return actual, err
		}

		// Retry if conflicted when align function is provided.
		if uo.AlignFunc != nil {
			return UpdateWithCtrlClient(ctx, cli, expected, opts...)
		}
	}

	return updated, err
}

type (
	DeleteClient interface {
		Delete(ctx context.Context, name string, opts meta.DeleteOptions) error
	}

	_DeleteOptions struct {
		meta.DeleteOptions
	}

	DeleteOption func(*_DeleteOptions)
)

// WithDeleteMetaOptions sets the delete options.
func WithDeleteMetaOptions(opts meta.DeleteOptions) DeleteOption {
	return func(do *_DeleteOptions) {
		do.DeleteOptions = opts
	}
}

// Delete will delete the resource if it exists.
//
// Delete doesn't return error if the resource is not found.
func Delete(ctx context.Context, cli DeleteClient, expected MetaObject, opts ...DeleteOption) error {
	if reflect.ValueOf(expected).IsZero() {
		return errors.New("expected is nil")
	}

	name := expected.GetName()
	if name == "" {
		return errors.New("resource name may not be empty")
	}

	var do _DeleteOptions
	for i := range opts {
		opts[i](&do)
	}

	err := cli.Delete(ctx, name, do.DeleteOptions)
	if err != nil && !kerrors.IsNotFound(err) {
		return err
	}

	return nil
}

// DeleteWithCtrlClient is similar to Delete, but uses the ctrl client.
func DeleteWithCtrlClient(ctx context.Context, cli ctrlcli.Client, expected MetaObject, opts ...DeleteOption) error {
	if reflect.ValueOf(expected).IsZero() {
		return errors.New("expected is nil")
	}

	name := expected.GetName()
	if name == "" {
		return errors.New("resource name may not be empty")
	}

	var do _DeleteOptions
	for i := range opts {
		opts[i](&do)
	}

	err := cli.Delete(ctx, expected, &ctrlcli.DeleteOptions{
		GracePeriodSeconds: do.GracePeriodSeconds,
		Preconditions:      do.Preconditions,
		PropagationPolicy:  do.PropagationPolicy,
		DryRun:             do.DryRun,
		Raw:                ptr.To(do.DeleteOptions),
	})
	if err != nil && !kerrors.IsNotFound(err) {
		return err
	}

	return nil
}
