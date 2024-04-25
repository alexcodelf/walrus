package kubehelper

import (
	"fmt"
	"time"
)

func NormalizeResourceHookName(resName string) string {
	return fmt.Sprintf("%s-hook", resName)
}

func NormalizeResourceRunName(resName string) string {
	return fmt.Sprintf("%s-run-%d", resName, time.Now().Unix())
}
