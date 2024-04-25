package parse

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/seal-io/utils/json"
	"k8s.io/apimachinery/pkg/util/sets"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourceruns/api"
	"github.com/seal-io/walrus/pkg/resources/interpolation"
)

const (
	// _variablePrefix the prefix of the variable name.
	_variablePrefix = "_walrus_var_"

	// _resourcePrefix the prefix of the resource output name.
	_resourcePrefix = "_walrus_res_"
)

// _interpolationReg is the regular expression for matching non-reference or non-variable expressions.
// Reference: https://developer.hashicorp.com/terraform/language/expressions/strings#escape-sequences-1
// To handle escape sequences, ${xxx} is converted to $${xxx}.
// If there are more than two consecutive $ symbols, like $${xxx}, they are further converted to $$${xxx}.
// During Terraform processing, $${} is ultimately transformed back to ${};
// this interpolation is used to ensure a WYSIWYG user experience.
var _interpolationReg = regexp.MustCompile(`\$\{((var\.)?([^.}]+)(?:\.([^.}]+))?)[^\}]*\}`)

// ParseModuleAttributes parse module variables and dependencies.
func ParseModuleAttributes(
	ctx context.Context,
	run *walruscore.ResourceRun,
	attributes map[string]any,
	onlyValidated bool,
) (attrs map[string]any, variables walrus.Variables, outputs map[string]api.OutputValue, err error) {
	var dependencyResourceOutputs []string

	replaced := !onlyValidated

	attrs, _, dependencyResourceOutputs, err = parseAttributeReplace(attributes, replaced)
	if err != nil {
		return
	}

	// TODO (alex): get the variables.

	variables = make(walrus.Variables, 0)

	if !onlyValidated {
		dependOutputMap := toDependOutputMap(dependencyResourceOutputs)

		outputs, err = getResourceDependencyOutputsByID(
			ctx,
			run.Spec.ResourceName,
			dependOutputMap)
		if err != nil {
			return nil, nil, nil, err
		}

		// Check if all dependency resource outputs are found.
		for outputName := range dependOutputMap {
			if _, ok := outputs[outputName]; !ok {
				return nil, nil, nil, fmt.Errorf("resource %s dependency output %s not found",
					run.Spec.ResourceName, outputName)
			}
		}
	}

	return attrs, variables, outputs, nil
}

// toDependOutputMap splits the dependencyResourceOutputs from {resource}_{resource_name}_{output_name}
// to a map of {resource_name}_{output_name}:{resource}.
func toDependOutputMap(dependencyResourceOutputs []string) map[string]string {
	dependOutputMap := make(map[string]string, 0)

	for _, dependOutput := range dependencyResourceOutputs {
		ss := strings.SplitN(dependOutput, "_", 2)
		if len(ss) != 2 {
			continue
		}
		dependOutputMap[ss[1]] = ss[0]
	}

	return dependOutputMap
}

// parseAttributeReplace parses attribute variable ${var.name} replaces it with ${var._variablePrefix+name},
// resource reference ${res.name.output} replaces it with ${var._resourcePrefix+name}
// and returns variable names and output names.
// Replaced flag indicates whether to replace the module attribute variable with prefix string.
func parseAttributeReplace(
	attributes map[string]any,
	replaced bool,
) (map[string]any, []string, []string, error) {
	bs, err := json.Marshal(attributes)
	if err != nil {
		return nil, nil, nil, err
	}

	variableMatches := interpolation.VariableReg.FindAllSubmatch(bs, -1)
	resourceMatches := interpolation.ResourceReg.FindAllSubmatch(bs, -1)

	variableMatched := sets.NewString()

	for _, match := range variableMatches {
		if len(match) > 1 {
			variableMatched.Insert(string(match[1]))
		}
	}

	resourceMatched := sets.NewString()

	for _, match := range resourceMatches {
		if len(match) > 1 {
			// Resource outputs are in the form:
			// - res_{resource_name}_{output_name}.
			resourceMatched.Insert("res_" + string(match[1]) + "_" + string(match[2]))
		}
	}

	variableRepl := "${var." + _variablePrefix + "${1}}"
	bs = interpolation.VariableReg.ReplaceAll(bs, []byte(variableRepl))

	resourceRepl := "${var." + _resourcePrefix + "${1}_${2}}"
	bs = interpolation.ResourceReg.ReplaceAll(bs, []byte(resourceRepl))

	// Replace interpolation from ${} to $${} to avoid escape sequences.
	bs = _interpolationReg.ReplaceAllFunc(bs, func(match []byte) []byte {
		m := _interpolationReg.FindSubmatch(match)

		if len(m) != 5 {
			return match
		}

		// If it is a variable or resource reference, do not replace.
		if string(m[2]) == "var." {
			return match
		}

		var b bytes.Buffer

		b.WriteString("$")
		b.Write(match)

		return b.Bytes()
	})

	if replaced {
		err = json.Unmarshal(bs, &attributes)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return attributes, variableMatched.List(), resourceMatched.List(), nil
}

// getResourceDependencyOutputsByID gets the dependency outputs of the resource by resource id.
func getResourceDependencyOutputsByID(
	ctx context.Context,
	resourceName string,
	dependOutputs map[string]string,
) (map[string]api.OutputValue, error) {
	// TODO (alex): get the dependency outputs of the resource by resource name.
	return nil, nil
}
