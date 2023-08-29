package server

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"k8s.io/client-go/rest"

	"github.com/seal-io/walrus/pkg/cache"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/utils/strs"
)

type initOptions struct {
	K8sConfig      *rest.Config
	K8sCacheReady  chan struct{}
	ModelClient    *model.Client
	SkipTLSVerify  bool
	DatabaseDriver *sql.DB
	CacheDriver    cache.Driver
}

func (r *Server) init(ctx context.Context, opts initOptions) error {
	// Initialize data for system.
 inits := []initiation{
 	r.applyModelSchemas,
 	r.setupSettings,
 	r.initConfigs,
 	r.registerMetricCollectors,
 	r.registerHealthCheckers,
 	r.startBackgroundJobs,
 	r.setupBusSubscribers,
 	r.initCatalog,
 }
	if r.EnableAuthn {
		inits = append(inits,
			r.configureCasdoor,
		)
	}

	// Initialize data for user.
	inits = append(inits,
		r.createBuiltinRbac,
		r.createBuiltinCatalogs,
		r.createBuiltinPerspectives,
		r.createBuiltinProjects,
	)

	for i := range inits {
		if err := inits[i](ctx, opts); err != nil {
			return fmt.Errorf("failed to %s: %w",
				loadInitiationName(inits[i]), err)
		}
	}

	return nil
}

type initiation func(context.Context, initOptions) error

func (r *Server) initCatalog(ctx context.Context, opts initOptions) error {
	// Initialize catalog from settings
	// If the setting does not exist or is empty, use the built-in catalog
	return nil
}

func loadInitiationName(i initiation) string {
	n := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	n = strings.TrimPrefix(strings.TrimSuffix(filepath.Ext(n), "-fm"), ".")
	return strs.Decamelize(n, true)
}
