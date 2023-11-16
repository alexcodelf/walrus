package runtimex

import (
	"runtime"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/seal-io/walrus/utils/log"
)

func init() {
	if strings.Contains(os.Args[0], "cli") {
		_, _ = maxprocs.Set(maxprocs.Logger(nil))
	} else {
		_, _ = maxprocs.Set(maxprocs.Logger(log.WithName("maxprocs").Printf))
	}
}

func NumCPU() int {
	return runtime.GOMAXPROCS(0)
}

import (
	"os"
	"runtime"
	"strings"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/seal-io/walrus/utils/log"
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(log.WithName("maxprocs").Printf))
}

func NumCPU() int {
	return runtime.GOMAXPROCS(0)
}
