package cron

import (
	"context"
	"errors"
	"fmt"
	"sync"

	settingbus "github.com/seal-io/seal/pkg/bus/setting"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/settings"
	"github.com/seal-io/seal/utils/cron"
	"github.com/seal-io/seal/utils/log"
)

type (
	// JobCreator is the creator for creating {cron.Expr, cron.Task} tuple,
	// the life of given context.Context ends by this creation,
	// do not use the long-term processing with this context.Context.
	JobCreator func(ctx context.Context, name, expr string) (Expr, Task, error)

	// JobCreators holds JobCreator with its name.
	JobCreators map[string]JobCreator
)

var (
	js = JobCreators{}
	o  sync.Once
)

// Register executes all job creators and
// schedules the returning task with the returning expression.
func Register(ctx context.Context, mc model.ClientSet, cs JobCreators) (err error) {
	err = errors.New("not allowed duplicated registering")
	o.Do(func() {
		for n, c := range cs {
			js[n] = c
		}
		err = doRegister(ctx, mc)
	})
	return
}

func doRegister(ctx context.Context, mc model.ClientSet) error {
	// NB(thxCode): don't stop the core cron scheduler.
	var err = cron.Start(ctx)
	if err != nil {
		return err
	}

	for n, c := range js {
		if c == nil {
			continue
		}

		var s = settings.Index(n)
		if s == nil {
			continue
		}
		// get cron expr of the job from global model client.
		var v string
		v, err = s.Value(ctx, mc)
		if err != nil {
			return fmt.Errorf("error gettting job cron expr: %w", err)
		}

		ce, ct, err := c(ctx, n, v)
		if err != nil {
			return fmt.Errorf("error creating %s job: %w", n, err)
		}
		err = cron.Schedule(n, ce, ct)
		if err != nil {
			return fmt.Errorf("error scheduling %s job: %w", n, err)
		}
	}
	return nil
}

// Sync observes the cron expr setting changes and re-register jobs.
func Sync(ctx context.Context, m settingbus.BusMessage) error {
	var logger = log.WithName("cronjobs")

	type job struct {
		Name string
		Expr Expr
		Task Task
	}

	var jobs []job
	for i := 0; i < len(m.Refer); i++ {
		if m.Refer[i] == nil {
			continue
		}

		var n = m.Refer[i].Name
		var c, exist = js[n]
		if !exist {
			continue
		}

		var s = settings.Index(n)
		if s == nil {
			continue
		}
		// get cron expr of the job from transactional model client.
		var v, err = s.Value(ctx, m.ModelClient)
		if err != nil {
			return fmt.Errorf("error gettting job cron expr: %w", err)
		}

		var j = job{Name: n}
		j.Expr, j.Task, err = c(ctx, n, v)
		if err != nil {
			return fmt.Errorf("error creating %s job: %w", n, err)
		}
		jobs = append(jobs, j)
	}

	for i := 0; i < len(jobs); i++ {
		var j = jobs[i]
		var err = cron.Schedule(j.Name, j.Expr, j.Task)
		if err != nil {
			// NB(thxCode): raising error cannot roll back successfully scheduled job in the same for-loop,
			// so just warn out here.
			logger.Errorf("error scheduling %s job: %v", j.Name, err)
		}
		// TODO(thxCode): support rolling back successfully scheduled job.
	}
	return nil
}