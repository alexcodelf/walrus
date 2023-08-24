package terraform

import (
	"context"
	"time"

	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/terraform/parser"
	"github.com/seal-io/walrus/utils/strs"
)

// updateServiceDriftResult update service and its drift result.
func updateServiceDriftResult(
	ctx context.Context,
	mc model.ClientSet,
	entity *model.Service,
	sr *model.ServiceRevision,
) error {
	if sr.Type != types.ServiceRevisionTypeDetect {
		return nil
	}

	serviceDriftResult := &types.ServiceDriftResult{
		Drifted: false,
		Time:    time.Now(),
	}

	// Parse service drift output.
	sdo, err := parser.ParseDriftOutput(sr.StatusMessage)
	if err != nil {
		return err
	}

	if sdo != nil && len(sdo.ResourceDrifts) > 0 {
		rds := make(map[string]*types.ResourceDrift, len(sdo.ResourceDrifts))
		for _, rd := range sdo.ResourceDrifts {
			rds[strs.Join("/", rd.Type, rd.Name)] = rd
		}

		if err = updateResourceDriftResult(ctx, mc, entity.ID, rds); err != nil {
			return err
		}

		serviceDriftResult.Drifted = true
		serviceDriftResult.Time = time.Now()
		serviceDriftResult.Result = sdo
	}

	if !serviceDriftResult.Drifted {
		if err := resetResourceDriftResult(ctx, mc, entity.ID); err != nil {
			return err
		}
	}

	return mc.Services().UpdateOne(entity).
		SetDriftResult(entity.DriftResult).
		Exec(ctx)
}

// updateResourceDriftResult update the drift detection result of the service's resources.
func updateResourceDriftResult(
	ctx context.Context,
	mc model.ClientSet,
	serviceID object.ID,
	resourceDrifts map[string]*types.ResourceDrift,
) error {
	resources, err := mc.ServiceResources().Query().
		Where(serviceresource.ServiceID(serviceID)).
		All(ctx)
	if err != nil && !model.IsNotFound(err) {
		return err
	}

	for i := range resources {
		r := resources[i]

		key := strs.Join("/", r.Type, r.Name)
		if _, ok := resourceDrifts[key]; !ok {
			continue
		}

		r.DriftResult = &types.ServiceResourceDriftResult{
			Drifted: true,
			Time:    time.Now(),
			Result:  resourceDrifts[key],
		}
	}

	return updateResources(ctx, mc, resources)
}

// resetResourceDriftResult reset the drift detection result of the service's resources.
func resetResourceDriftResult(ctx context.Context, mc model.ClientSet, serviceID object.ID) error {
	resources, err := mc.ServiceResources().Query().
		Where(serviceresource.ServiceID(serviceID)).
		All(ctx)
	if err != nil && !model.IsNotFound(err) {
		return err
	}

	for _, resource := range resources {
		resource.DriftResult = nil
	}

	return updateResources(ctx, mc, resources)
}

func updateResources(ctx context.Context, mc model.ClientSet, resources model.ServiceResources) error {
	updates, err := dao.ServiceResourceUpdates(mc, resources...)
	if err != nil {
		return err
	}

	for _, update := range updates {
		err = update.Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
