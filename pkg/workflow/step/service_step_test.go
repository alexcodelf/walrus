package step

import (
	"fmt"
	"testing"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/zclconf/go-cty/cty"
)

var ss = &model.WorkflowStep{
	Name:        "test",
	Description: "Description of test",
	Type:        "test",
	Annotations: map[string]string{
		"seal.io/workflow": "test",
	},
	Timeout: 0,
	Spec: map[string]interface{}{
		"image": "docker/whalesay",
		"Command": []string{
			"cowsay",
			"hello world",
		},
	},
}

var testTemplate = &model.WorkflowStepTemplate{
	Name: "walrus-service",
	Schema: property.Schemas{
		{
			Name:        "serviceName",
			Type:        cty.String,
			Description: "name of the service",
		},
		{
			Name:        "template",
			Type:        cty.String,
			Description: "template of the service",
		},
		{
			Name:        "templateVersion",
			Type:        cty.String,
			Description: "version of the template",
		},
	},
}

func TestConvertStepToTemplate(t *testing.T) {
	fmt.Println(ss)
}
