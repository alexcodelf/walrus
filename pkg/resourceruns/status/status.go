package status

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/system"
)

const (
	ActionDelete = "delete"
	ActionStop   = "stop"
)

// IsStatusRunning checks if the resource run is in the running status.
func IsStatusRunning(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionRunning.IsUnknown(run)
}

func IsStatusPending(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionPending.IsUnknown(run)
}

func IsStatusFailed(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionPending.IsFalse(run) ||
		apistatus.ResourceRunConditionRunning.IsFalse(run)
}

func IsStatusPlanned(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionRunning.IsTrue(run)
}

func IsStatusSucceeded(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionRunning.IsTrue(run)
}

func IsStatusCanceled(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionCanceled.IsTrue(run)
}

func IsPlanFailed(run *walruscore.ResourceRun) bool {
	return apistatus.ResourceRunConditionRunning.IsFalse(run)
}

// SetStatusFalse sets the status of the resource run to false.
func SetStatusFalse(run *walruscore.ResourceRun, errMsg string) {
	logger := klog.Background().WithName("resource-run").WithName("status")

	switch {
	case apistatus.ResourceRunConditionRunning.IsUnknown(run):
		errMsg = fmt.Sprintf("plan failed: %s", errMsg)
		apistatus.ResourceRunConditionRunning.False(run, apistatus.ResourceRunConditionReasonFailed, errMsg)
	case apistatus.ResourceRunConditionPending.IsUnknown(run):
		errMsg = fmt.Sprintf("pending failed: %s", errMsg)
		apistatus.ResourceRunConditionPending.False(run, apistatus.ResourceRunConditionReasonFailed, errMsg)
	default:
		logger.Info("cannot set status false, unknown status: %s", run.Status.Phase)
		return
	}

	run.Status.ConditionSummary = *apistatus.WalkResourceRun(&run.Status.StatusDescriptor)
}

// SetStatusTrue sets the status of the resource run to true.
// It marks the status of the resource run as "Succeeded".
func SetStatusTrue(run *walruscore.ResourceRun, msg string) {
	logger := klog.Background().WithName("resource-run").WithName("status")

	switch {
	case apistatus.ResourceRunConditionRunning.IsUnknown(run):
		apistatus.ResourceRunConditionRunning.True(run, apistatus.ResourceRunConditionReasonSucceeded, "")
	default:
		logger.Infof("cannot set status true, unknown status: %s", run.Status.Phase)
		return
	}

	run.Status.ConditionSummary = *apistatus.WalkResourceRun(&run.Status.StatusDescriptor)
}

// UpdateStatus updates the status of the resource run.
func UpdateStatus(ctx context.Context, obj *walruscore.ResourceRun) (*walruscore.ResourceRun, error) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	// Update resource run status.
	obj.Status.ConditionSummary = *apistatus.WalkResourceRun(&obj.Status.StatusDescriptor)
	obj, err := loopbackKubeClient.WalruscoreV1().ResourceRuns(obj.Namespace).UpdateStatus(ctx, obj, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// UpdateStatusWithErr updates the status of the resource run with the given error.
func UpdateStatusWithErr(ctx context.Context, run *walruscore.ResourceRun, err error) (*walruscore.ResourceRun, error) {
	if err == nil {
		return run, nil
	}

	SetStatusFalse(run, err.Error())

	return UpdateStatus(ctx, run)
}

// CheckDependencyStatus checks the resource dependency status of the resource run.
// Works for create, update, rollback and start actions.
func CheckDependencyStatus(
	ctx context.Context,
	run *walruscore.ResourceRun,
) (bool, error) {
	// Check dependencies.
	// TODO (alex): check the dependencies of the resource run.

	return true, nil
}

// CheckDependantStatus checks the resource dependant status of the resource run.
// Works for both stop and delete actions.
func CheckDependantStatus(
	ctx context.Context,
	run *walruscore.ResourceRun,
) (bool, error) {
	// Check dependants.
	// TODO (alex): check the dependants of the resource run.

	return true, nil
}

// CheckStatus checks the status of the resource run dependant or dependency with the given resource run.
func CheckStatus(ctx context.Context, run *walruscore.ResourceRun) (bool, error) {
	// Check the status of the resource run.
	// TODO (alex): check the status of the resource run.

	return true, nil
}
