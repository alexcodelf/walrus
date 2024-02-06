package terraform

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/servervars"
)

const (
	// _applyCommands the commands to apply deployment of the resource run.
	_applyCommands = "terraform init -no-color && terraform apply -auto-approve -no-color"
	// _destroyCommands the commands to destroy deployment of the resource run.
	_destroyCommands = "terraform init -no-color && terraform destroy -auto-approve -no-color"

	// _planFileName the file to store the plan of the resource run.
	_planFileName = "plan.json"
	// _planAPI.
	_planAPI = "/v1/projects/%s/environments/%s/resources/%s/runs/%s/component-change"

	// _detectCommands the commands to detect drift of the revision.
	_planCommands = "terraform init -no-color && terraform plan -no-color -out=plan.out %s" +
		" && TF_LOG=ERROR terraform show -json plan.out > " + _planFileName
)

func getPlanCommands(run *model.ResourceRun, varFile string, opts JobCreateOptions) string {
	planCommands := fmt.Sprintf(_planCommands, varFile)

	planAPI := fmt.Sprintf("%s%s", opts.ServerURL,
		fmt.Sprintf(_planAPI,
			run.ProjectID,
			run.EnvironmentID,
			run.ResourceID,
			run.ID))

	planCommands += fmt.Sprintf(
		" && curl -s -f -X PUT -H \"Content-Type: application/json\" -H \"Authorization: Bearer $ACCESS_TOKEN\" %s -d @%s",
		planAPI,
		_planFileName,
	)

	if !servervars.TlsCertified.Get() {
		planCommands += " -k"
	}

	return planCommands
}

func getApplyCommands() string {
	return _applyCommands
}

func getDestroyCommands() string {
	return _destroyCommands
}
