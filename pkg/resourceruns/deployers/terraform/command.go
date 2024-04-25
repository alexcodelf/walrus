package terraform

import (
	"fmt"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourceruns/deployer"
	servervars "github.com/seal-io/walrus/pkg/servers/vars"
)

const (
	// _planFileName the file to store the plan of the resource run.
	_planFileName = "plan.out"
	// _jsonPlanFileName the json file to show the plan of the resource run.
	_jsonPlanFileName = "plan.json"

	// _planCommands the commands to get the changes of the resource run.
	_planCommands = "terraform init -no-color && terraform plan %s -no-color -out=plan.out %s" +
		" && terraform show -json plan.out > " + _jsonPlanFileName
	// _applyCommands the commands to apply deployment of the resource run.
	_applyCommands = "terraform init -no-color && terraform apply %s -no-color"
	// _destroyCommands the commands to destroy deployment of the resource run.
	// As destroy planned in plan file, use apply command to execution the plan.
	_destroyCommands = "terraform init -no-color && terraform apply %s -no-color"

	// _planAPI.
	_planAPI = "/v1/projects/%s/environments/%s/resources/%s/runs/%s/plan"
)

func getPlanCommands(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) string {
	var (
		destroy string
		varfile = fmt.Sprintf(" -var-file=%s/terraform.tfvars", _secretMountPath)
	)

	if run.Spec.Type == walruscore.ResourceRunTypeDelete || run.Spec.Type == walruscore.ResourceRunTypeStop {
		destroy = "-destroy"
	}

	return fmt.Sprintf(_planCommands, destroy, varfile) + setPlanFile(run, opts)
}

func getApplyCommands(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) string {
	return fmt.Sprintf("%s && %s", getPlanFile(run, opts), fmt.Sprintf(_applyCommands, _planFileName))
}

func getDestroyCommands(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) string {
	return fmt.Sprintf("%s && %s", getPlanFile(run, opts), fmt.Sprintf(_destroyCommands, _planFileName))
}

// getPlanFile returns the command to get the plan file.
func getPlanFile(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) string {
	getPlanAPI := fmt.Sprintf("%s%s", opts.ServerURL,
		fmt.Sprintf(_planAPI, run.Spec.Project, run.Namespace, run.Spec.ResourceName, run.Name))

	getPlan := fmt.Sprintf(
		"curl -sS --fail-with-body -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" %s -o %s",
		getPlanAPI,
		_planFileName,
	)

	if !servervars.IsTlsCertified() {
		getPlan += " -k"
	}

	return getPlan
}

// setPlanFile returns the command to set the plan file.
func setPlanFile(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) string {
	setPlanAPI := fmt.Sprintf("%s%s", opts.ServerURL,
		fmt.Sprintf(_planAPI,
			run.Spec.Project,
			run.Namespace,
			run.Spec.ResourceName,
			run.Name))

	setPlan := fmt.Sprintf(
		" && curl -sS --fail-with-body -X POST -H \"Content-Type: multipart/form-data\" -H \"Authorization: Bearer $ACCESS_TOKEN\""+
			" %s -F jsonplan=@%s -F plan=@%s",
		setPlanAPI,
		_jsonPlanFileName,
		_planFileName,
	)

	if !servervars.IsTlsCertified() {
		setPlan += " -k"
	}

	return setPlan
}
