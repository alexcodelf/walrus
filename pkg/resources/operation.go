package resources

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponent"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	deptypes "github.com/seal-io/walrus/pkg/deployer/types"
	pkgrun "github.com/seal-io/walrus/pkg/resourceruns"
	runjob "github.com/seal-io/walrus/pkg/resourceruns/job"
	runstatus "github.com/seal-io/walrus/pkg/resourceruns/status"
	resstatus "github.com/seal-io/walrus/pkg/resources/status"
	"github.com/seal-io/walrus/utils/errorx"
	"github.com/seal-io/walrus/utils/log"
)

const (
	ActionDelete = "delete"
	ActionStop   = "stop"
)

// Options for deploy or destroy.
type Options struct {
	// Deployer The deployer type to perform the resource run.
	Deployer deptypes.Deployer

	// Draft if resource is in draft status.
	Draft bool

	// RunType The type of the resource run.
	RunType types.RunType
	// Preview if the resource run requires preview.
	Preview bool
	// ChangeComment of the resource run.
	ChangeComment string
}

// Basic Operations.

func Create(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Resource,
	opts Options,
) (res *model.Resource, run *model.ResourceRun, err error) {
	if err = SetSubjectID(ctx, entity); err != nil {
		return nil, nil, err
	}

	if err = SetDefaultLabels(ctx, mc, entity); err != nil {
		return nil, nil, err
	}
	entity.IsModified = true

	if opts.Draft {
		status.ResourceStatusUnDeployed.True(entity, "Draft")
		entity.Status.SetSummary(status.WalkResource(&entity.Status))

		entity, err = mc.Resources().Create().
			Set(entity).
			SaveE(ctx, dao.ResourceDependenciesEdgeSave)
		if err != nil {
			return nil, nil, err
		}

		return entity, nil, nil
	}

	// Save the resource.
	err = mc.WithTx(ctx, func(tx *model.Tx) (err error) {
		// TODO(thxCode): generated by entc.

		status.ResourceStatusUnDeployed.True(entity, "")
		entity.Status.SetSummary(status.WalkResource(&entity.Status))

		entity, err = tx.Resources().Create().
			Set(entity).
			SaveE(ctx, dao.ResourceDependenciesEdgeSave)
		if err != nil {
			return err
		}

		// Create resource state.
		err = tx.ResourceStates().Create().
			SetResourceID(entity.ID).
			SetData("").
			Exec(ctx)
		if err != nil {
			return err
		}

		// Create resource run.
		run, err = pkgrun.Create(ctx, tx, pkgrun.CreateOptions{
			ResourceID:    entity.ID,
			DeployerType:  opts.Deployer.Type(),
			Type:          types.RunTypeCreate,
			ChangeComment: opts.ChangeComment,
			Preview:       opts.Preview,
		})

		return err
	})
	if err != nil {
		return nil, nil, err
	}

	defer errorHandler(mc, entity, run, status.ResourceStatusDeployed, err)

	// Check dependency status.
	ready, err := CheckDependencyStatus(ctx, mc, opts.Deployer, entity)
	if err != nil {
		return nil, nil, err
	}

	if !ready {
		return entity, run, nil
	}

	err = runjob.PerformRunJob(ctx, mc, opts.Deployer, run)

	return entity, run, err
}

// Upgrade upgrades the resource.
func Upgrade(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Resource,
	opts Options,
) (*model.ResourceRun, error) {
	opts.RunType = types.RunTypeUpdate
	return upgrade(ctx, mc, entity, opts)
}

type DeleteOptions struct {
	Options

	WithoutCleanup bool
}

func Delete(ctx context.Context, mc model.ClientSet, entity *model.Resource, opts DeleteOptions) (err error) {
	var run *model.ResourceRun

	err = SetSubjectID(ctx, entity)
	if err != nil {
		return err
	}

	// If no resource component exists, skip calling deployer destroy and do straight deletion.
	exist, err := mc.ResourceComponents().Query().
		Where(resourcecomponent.ResourceID(entity.ID)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if !exist || opts.WithoutCleanup {
		return mc.Resources().DeleteOneID(entity.ID).Exec(ctx)
	}

	if status.ResourceStatusDeployed.IsUnknown(entity) {
		return fmt.Errorf("cannot delete resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	entity, err = mc.Resources().UpdateOne(entity).
		Set(entity).
		Save(ctx)
	if err != nil {
		return err
	}

	defer errorHandler(mc, entity, run, status.ResourceStatusDeleted, err)

	run, err = pkgrun.Create(ctx, mc, pkgrun.CreateOptions{
		ResourceID:    entity.ID,
		DeployerType:  opts.Deployer.Type(),
		Type:          types.RunTypeDelete,
		ChangeComment: opts.ChangeComment,
		Preview:       opts.Preview,
	})
	if err != nil {
		return err
	}

	// Check dependant status.
	ready, err := CheckDependantStatus(ctx, mc, entity, ActionStop)
	if err != nil {
		return err
	}

	if !ready {
		return nil
	}

	err = runjob.PerformRunJob(ctx, mc, opts.Deployer, run)

	return err
}

// Start starts the resource.
func Start(ctx context.Context, mc model.ClientSet, entity *model.Resource, opts Options) (*model.ResourceRun, error) {
	if !resstatus.IsInactive(entity) {
		return nil, fmt.Errorf("cannot start resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	opts.RunType = types.RunTypeStart

	return upgrade(ctx, mc, entity, opts)
}

func upgrade(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Resource,
	opts Options,
) (*model.ResourceRun, error) {
	var run *model.ResourceRun

	err := SetSubjectID(ctx, entity)
	if err != nil {
		return nil, err
	}

	if status.ResourceStatusProgressing.IsUnknown(entity) {
		return nil, fmt.Errorf("cannot upgrade resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	entity.IsModified = true

	// Update Status and ave the resource.
	if opts.Draft {
		err := mc.Resources().UpdateOne(entity).
			Set(entity).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	if err := SetSubjectID(ctx, entity); err != nil {
		return nil, err
	}

	entity, err = mc.Resources().UpdateOne(entity).
		Set(entity).
		SaveE(ctx, dao.ResourceDependenciesEdgeSave)
	if err != nil {
		return nil, err
	}

	defer errorHandler(mc, entity, run, status.ResourceStatusDeployed, err)

	run, err = pkgrun.Create(ctx, mc, pkgrun.CreateOptions{
		ResourceID:    entity.ID,
		DeployerType:  opts.Deployer.Type(),
		Type:          opts.RunType,
		ChangeComment: opts.ChangeComment,
		Preview:       opts.Preview,
	})
	if err != nil {
		return nil, err
	}

	ready, err := CheckDependencyStatus(ctx, mc, opts.Deployer, entity)
	if err != nil {
		return nil, errorx.Wrap(err, "error checking dependency status")
	}

	if !ready {
		return run, nil
	}

	// Perform the resource run.
	err = runjob.PerformRunJob(ctx, mc, opts.Deployer, run)

	return run, err
}

func Stop(ctx context.Context, mc model.ClientSet, entity *model.Resource, opts Options) (err error) {
	if !IsStoppable(entity) {
		return fmt.Errorf("resource %s is non-stoppable", entity.Name)
	}

	if !CanBeStopped(entity) {
		return fmt.Errorf("cannot stop resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	if status.ResourceStatusProgressing.IsUnknown(entity) {
		return fmt.Errorf("cannot stop resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	err = SetSubjectID(ctx, entity)
	if err != nil {
		return err
	}

	// If no resource component exists, skip calling deployer destroy and do straight deletion.
	exist, err := mc.ResourceComponents().Query().
		Where(resourcecomponent.ResourceID(entity.ID)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if !exist {
		return mc.Resources().DeleteOneID(entity.ID).Exec(ctx)
	}

	var run *model.ResourceRun

	entity.IsModified = true
	err = mc.Resources().UpdateOne(entity).
		Set(entity).
		Exec(ctx)
	if err != nil {
		return err
	}

	defer errorHandler(mc, entity, run, status.ResourceStatusStopped, err)

	run, err = pkgrun.Create(ctx, mc, pkgrun.CreateOptions{
		ResourceID:    entity.ID,
		DeployerType:  opts.Deployer.Type(),
		Type:          types.RunTypeStop,
		ChangeComment: opts.ChangeComment,
		Preview:       opts.Preview,
	})
	if err != nil {
		return err
	}

	// Check dependant status.
	ready, err := CheckDependantStatus(ctx, mc, entity, ActionDelete)
	if err != nil {
		return err
	}

	if !ready {
		return nil
	}

	err = runjob.PerformRunJob(ctx, mc, opts.Deployer, run)

	return err
}

func Rollback(ctx context.Context, mc model.ClientSet, resourceID, runID object.ID, opts Options) error {
	run, err := mc.ResourceRuns().Query().
		Where(
			resourcerun.ID(runID),
			resourcerun.ResourceID(resourceID),
		).
		WithResource().
		Only(ctx)
	if err != nil {
		return err
	}

	switch run.Type {
	case types.RunTypeCreate.String(), types.RunTypeUpdate.String():
	default:
		return errorx.Errorf("cannot rollback resource run %s: not a create or upgrade run", run.ID)
	}

	entity := run.Edges.Resource
	entity.Attributes = run.Attributes
	entity.ComputedAttributes = run.ComputedAttributes
	entity.IsModified = true

	if status.ResourceStatusProgressing.IsUnknown(entity) {
		return fmt.Errorf("cannot rollback resource %q: in %q status", entity.Name, entity.Status.SummaryStatus)
	}

	if entity.TemplateID != nil {
		// Find a previous template version when the resource is using template not definition.
		tv, err := mc.TemplateVersions().Query().
			Where(
				templateversion.Version(run.TemplateVersion),
				templateversion.TemplateID(run.TemplateID)).
			Only(ctx)
		if err != nil {
			return err
		}

		entity.TemplateID = &tv.ID
	}

	if err = SetSubjectID(ctx, entity); err != nil {
		return err
	}

	entity, err = mc.Resources().UpdateOne(entity).
		Set(entity).
		SaveE(ctx, dao.ResourceDependenciesEdgeSave)
	if err != nil {
		return errorx.Wrap(err, "error updating resource")
	}

	defer errorHandler(mc, entity, run, status.ResourceStatusDeployed, err)

	// Create resource run.
	run, err = pkgrun.Create(ctx, mc, pkgrun.CreateOptions{
		ResourceID:    entity.ID,
		DeployerType:  opts.Deployer.Type(),
		Type:          types.RunTypeRollback,
		ChangeComment: opts.ChangeComment,
		Preview:       opts.Preview,
	})
	if err != nil {
		return err
	}

	ready, err := CheckDependencyStatus(ctx, mc, opts.Deployer, entity)
	if err != nil {
		return err
	}

	if !ready {
		return nil
	}

	err = runjob.PerformRunJob(ctx, mc, opts.Deployer, run)

	return err
}

// Batch operations.

// CollectionUpgrade upgrades the resources.
func CollectionUpgrade(ctx context.Context, mc model.ClientSet, entities model.Resources, opts Options) error {
	// Group resources by environment.
	groupedResources, err := GroupByEnvironment(ctx, mc, entities)
	if err != nil {
		return err
	}

	opts.RunType = types.RunTypeUpdate

	for envID := range groupedResources {
		envResources, err := TopologicalSortResources(groupedResources[envID])
		if err != nil {
			return err
		}

		for _, entity := range envResources {
			_, err = upgrade(ctx, mc, entity, opts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectionDelete deletes the resources.
func CollectionDelete(ctx context.Context, mc model.ClientSet, entities model.Resources, opts DeleteOptions) error {
	// Group resources by environment.
	groupedResources, err := GroupByEnvironment(ctx, mc, entities)
	if err != nil {
		return err
	}

	for envID := range groupedResources {
		envResources, err := ReverseTopologicalSortResources(groupedResources[envID])
		if err != nil {
			return err
		}

		for _, entity := range envResources {
			err = Delete(ctx, mc, entity, opts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectionStart starts the resources.
func CollectionStart(ctx context.Context, mc model.ClientSet, entities model.Resources, opts Options) error {
	// Group resources by environment.
	groupedResources, err := GroupByEnvironment(ctx, mc, entities)
	if err != nil {
		return err
	}

	opts.RunType = types.RunTypeStart

	for envID := range groupedResources {
		envResources, err := TopologicalSortResources(groupedResources[envID])
		if err != nil {
			return err
		}

		for _, entity := range envResources {
			_, err = Start(ctx, mc, entity, opts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectionStop stops the resources.
func CollectionStop(ctx context.Context, mc model.ClientSet, entities model.Resources, opts Options) error {
	// Group resources by environment.
	groupedResources, err := GroupByEnvironment(ctx, mc, entities)
	if err != nil {
		return err
	}

	for envID := range groupedResources {
		envResources, err := ReverseTopologicalSortResources(groupedResources[envID])
		if err != nil {
			return err
		}

		for _, entity := range envResources {
			err = Stop(ctx, mc, entity, opts)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectionCreate creates the resources.
func CollectionCreate(
	ctx context.Context,
	mc model.ClientSet,
	entities model.Resources,
	opts Options,
) (model.Resources, error) {
	// Group resources by environment.
	groupedResources, err := GroupByEnvironment(ctx, mc, entities)
	if err != nil {
		return nil, err
	}

	createdResources := make(model.Resources, 0)

	for envID := range groupedResources {
		envResources, err := TopologicalSortResources(groupedResources[envID])
		if err != nil {
			return nil, err
		}

		for i := range envResources {
			entity := envResources[i]

			entity, _, err = Create(ctx, mc, entity, opts)
			if err != nil {
				return nil, err
			}

			createdResources = append(createdResources, entity)
		}
	}

	return createdResources, nil
}

// errorHandler handles the error of the resource operation.
func errorHandler(mc model.ClientSet, res *model.Resource, run *model.ResourceRun, ct status.ConditionType, err error) {
	if err == nil || res == nil {
		return
	}

	var (
		ctx    = context.Background()
		logger = log.WithName("resource")
	)

	ct.False(res, err.Error())

	updateErr := resstatus.UpdateStatus(ctx, mc, res)
	if updateErr != nil {
		logger.Errorf("error updating status of resource %s: %v", res.ID, err)
	}

	if run == nil {
		return
	}

	runstatus.SetStatusFalse(run, err.Error())

	_, updateErr = runstatus.UpdateStatus(ctx, mc, run)
	if updateErr != nil {
		logger.Errorf("error updating status of resource run %s: %v", run.ID, updateErr)
	}
}

// GroupByEnvironment groups the resources by environment.
func GroupByEnvironment(
	ctx context.Context,
	mc model.ClientSet,
	entities model.Resources,
) (map[object.ID]model.Resources, error) {
	err := PopulateEnvironmentID(ctx, mc, entities)
	if err != nil {
		return nil, err
	}

	groupedResources := make(map[object.ID]model.Resources)

	for _, entity := range entities {
		environmentID := entity.EnvironmentID

		if _, ok := groupedResources[environmentID]; !ok {
			groupedResources[environmentID] = make(model.Resources, 0)
		}

		groupedResources[environmentID] = append(groupedResources[environmentID], entity)
	}

	return groupedResources, nil
}

// PopulateEnvironmentID populate the environment id of the resources.
// As the environment id may not be included in the resource list returned by the query,
// we need to refill the environment id of the resources in some cases.
func PopulateEnvironmentID(ctx context.Context, mc model.ClientSet, entities model.Resources) error {
	var (
		logger = log.WithName("resource")
		ids    = make([]object.ID, 0, len(entities))
	)

	for _, entity := range entities {
		// Skip if the environment id is already set.
		if entity.EnvironmentID.Valid() {
			continue
		}

		ids = append(ids, entity.ID)
	}

	if len(ids) == 0 {
		return nil
	}

	fetchedResources, err := mc.Resources().Query().
		Select(resource.FieldID, resource.FieldEnvironmentID).
		Where(resource.IDIn(ids...)).
		All(ctx)
	if err != nil {
		return err
	}

	resourceEnvMap := make(map[object.ID]object.ID)
	for _, r := range fetchedResources {
		resourceEnvMap[r.ID] = r.EnvironmentID
	}

	for _, entity := range entities {
		if !entity.ID.Valid() {
			logger.Warnf("resource %s has no id", entity.Name)
			continue
		}

		if _, ok := resourceEnvMap[entity.ID]; !ok {
			return fmt.Errorf("resource %s has no environment id", entity.Name)
		}

		entity.EnvironmentID = resourceEnvMap[entity.ID]
	}

	return nil
}

// PerformResource get resource latest run and perform run job.
func PerformResource(ctx context.Context, mc model.ClientSet, dp deptypes.Deployer, resourceID object.ID) error {
	runs, err := dao.GetResourcesLatestRuns(ctx, mc, resourceID)
	if err != nil {
		return err
	}

	if len(runs) == 0 {
		return fmt.Errorf("latest resource run not found")
	}

	return runjob.PerformRunJob(ctx, mc, dp, runs[0])
}
