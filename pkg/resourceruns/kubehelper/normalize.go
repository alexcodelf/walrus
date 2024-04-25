package kubehelper

import (
	"fmt"
)

func NormalizeResourceRunConfigSecretName(resRunName string) string {
	return fmt.Sprintf("%s-config", resRunName)
}

func NormalizeResourceRunPlanWorkflowName(resRunName string) string {
	return fmt.Sprintf("%s-plan", resRunName)
}

func NormalizeResourceRunApproveApplyWorkflowName(resRunName string) string {
	return fmt.Sprintf("%s-approve-apply", resRunName)
}

func NormalizeResourceRunServiceAccountName(resRunName string) string {
	return fmt.Sprintf("%s-sa", resRunName)
}

func NormalizeResourceRunRoleBindingName(resRunName string) string {
	return fmt.Sprintf("%s-rb", resRunName)
}
