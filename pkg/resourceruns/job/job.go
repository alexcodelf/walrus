package job

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	deptypes "github.com/seal-io/walrus/pkg/deployer/types"
	"github.com/seal-io/walrus/pkg/resourceruns/status"
)

// PerformRunJob performs the run job by the given run.
// Depending on the run type and status, the deployer will perform different actions.
func PerformRunJob(ctx context.Context, mc model.ClientSet, dp deptypes.Deployer, run *model.ResourceRun) error {
	runJobType, err := GetRunJobType(run)
	if err != nil {
		return err
	}

	switch runJobType {
	case types.RunTaskTypePlan:
		return dp.Plan(ctx, mc, run, deptypes.PlanOptions{})
	case types.RunTaskTypeApply:
		return dp.Apply(ctx, mc, run, deptypes.ApplyOptions{})
	case types.RunTaskTypeDestroy:
		return dp.Destroy(ctx, mc, run, deptypes.DestroyOptions{})
	}

	return fmt.Errorf("unknown run job type %s", runJobType)
}

// GetRunJobType gets the run job type with its type and status.
// It makes the following decision.
//
//	| Run type         | Run status       | Job type         |
//	| ---------------- | ---------------- | ---------------- |
//	| create           | pending          | plan             |
//	| create           | planed           | apply            |
//	| upgrade          | pending          | plan             |
//	| upgrade          | planed           | apply            |
//	| delete           | pending          | plan             |
//	| delete           | planed           | destroy          |
//	| start            | pending          | plan             |
//	| start            | planed           | apply            |
//	| stop             | pending          | plan             |
//	| stop             | planed           | destroy          |
//	| rollback         | pending          | plan             |
//	| rollback         | planed           | apply            |
func GetRunJobType(run *model.ResourceRun) (types.RunJobType, error) {
	if status.IsStatusPending(run) {
		return types.RunTaskTypePlan, nil
	}

	if status.IsStatusPlanned(run) {
		switch types.RunType(run.Type) {
		case types.RunTypeCreate, types.RunTypeUpgrade, types.RunTypeStart, types.RunTypeRollback:
			return types.RunTaskTypeApply, nil
		case types.RunTypeDelete, types.RunTypeStop:
			return types.RunTaskTypeDestroy, nil
		}
	}

	return "", fmt.Errorf("unknown run type %s and status %s", run.Type, run.Status.SummaryStatus)
}
