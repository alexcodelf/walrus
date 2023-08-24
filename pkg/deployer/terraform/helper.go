package terraform

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/terraform/config"
	"github.com/seal-io/walrus/pkg/terraform/parser"
)

// parseAttributeReplace parses attribute variable ${var.name} replaces it with ${var._variablePrefix+name},
// service reference ${service.name.output} replaces it with ${var._servicePrefix+name}
// and returns variable names and service names.
func parseAttributeReplace(
	attributes map[string]any,
	variableNames []string,
	serviceOutputs []string,
) ([]string, []string) {
	for key, value := range attributes {
		if value == nil {
			continue
		}

		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			if _, ok := value.(map[string]any); !ok {
				continue
			}

			variableNames, serviceOutputs = parseAttributeReplace(
				value.(map[string]any),
				variableNames,
				serviceOutputs,
			)
		case reflect.String:
			str := value.(string)
			matches := _variableReg.FindAllStringSubmatch(str, -1)
			serviceMatches := _serviceReg.FindAllStringSubmatch(str, -1)

			var matched []string

			for _, match := range matches {
				if len(match) > 1 {
					matched = append(matched, match[1])
				}
			}

			var serviceMatched []string

			for _, match := range serviceMatches {
				if len(match) > 1 {
					serviceMatched = append(serviceMatched, match[1]+"_"+match[2])
				}
			}

			variableNames = append(variableNames, matched...)
			variableRepl := "${var." + _variablePrefix + "${1}}"
			str = _variableReg.ReplaceAllString(str, variableRepl)

			serviceOutputs = append(serviceOutputs, serviceMatched...)
			serviceRepl := "${var." + _servicePrefix + "${1}_${2}}"

			attributes[key] = _serviceReg.ReplaceAllString(str, serviceRepl)
		case reflect.Slice:
			if _, ok := value.([]any); !ok {
				continue
			}

			for _, v := range value.([]any) {
				if _, ok := v.(map[string]any); !ok {
					continue
				}
				variableNames, serviceOutputs = parseAttributeReplace(
					v.(map[string]any),
					variableNames,
					serviceOutputs,
				)
			}
		}
	}

	return variableNames, serviceOutputs
}

func getVarConfigOptions(variables model.Variables, serviceOutputs map[string]parser.OutputState) config.CreateOptions {
	varsConfigOpts := config.CreateOptions{
		Attributes: map[string]any{},
	}

	for _, v := range variables {
		varsConfigOpts.Attributes[_variablePrefix+v.Name] = v.Value
	}

	// Setup service outputs.
	for n, v := range serviceOutputs {
		varsConfigOpts.Attributes[_servicePrefix+n] = v.Value
	}

	return varsConfigOpts
}

func getModuleConfig(
	revision *model.ServiceRevision,
	template *model.TemplateVersion,
	opts createK8sSecretsOptions,
) (*config.ModuleConfig, error) {
	var (
		props              = make(property.Properties, len(revision.Attributes))
		typesWith          = revision.Attributes.TypesWith(template.Schema.Variables)
		sensitiveVariables = sets.Set[string]{}
	)

	for k, v := range revision.Attributes {
		props[k] = property.Property{
			Type:  typesWith[k],
			Value: v,
		}
	}

	attrs, err := props.TypedValues()
	if err != nil {
		return nil, err
	}

	mc := &config.ModuleConfig{
		Name:       opts.ServiceName,
		Source:     template.Source,
		Schema:     template.Schema,
		Attributes: attrs,
	}

	if template.Schema == nil {
		return mc, nil
	}

	for _, v := range template.Schema.Variables {
		// Add sensitive from schema variable.
		if v.Sensitive {
			sensitiveVariables.Insert(fmt.Sprintf(`var\.%s`, v.Name))
		}

		// Add seal metadata.
		var attrValue string

		switch v.Name {
		case WalrusMetadataProjectName:
			attrValue = opts.ProjectName
		case WalrusMetadataEnvironmentName:
			attrValue = opts.EnvironmentName
		case WalrusMetadataServiceName:
			attrValue = opts.ServiceName
		case WalrusMetadataProjectID:
			attrValue = opts.ProjectID.String()
		case WalrusMetadataEnvironmentID:
			attrValue = opts.EnvironmentID.String()
		case WalrusMetadataServiceID:
			attrValue = opts.ServiceID.String()
		case WalrusMetadataNamespaceName:
			attrValue = opts.ManagedNamespaceName
		}

		if attrValue != "" {
			mc.Attributes[v.Name] = attrValue
		}
	}

	sensitiveVariableRegex, err := matchAnyRegex(sensitiveVariables.UnsortedList())
	if err != nil {
		return nil, err
	}

	mc.Outputs = make([]config.Output, len(template.Schema.Outputs))
	for i, v := range template.Schema.Outputs {
		mc.Outputs[i].Sensitive = v.Sensitive
		mc.Outputs[i].Name = v.Name
		mc.Outputs[i].ServiceName = opts.ServiceName
		mc.Outputs[i].Value = v.Value

		if v.Sensitive {
			continue
		}

		// Update sensitive while output is from sensitive data, like secret.
		if sensitiveVariables.Len() != 0 && sensitiveVariableRegex.Match(v.Value) {
			mc.Outputs[i].Sensitive = true
		}
	}

	return mc, nil
}

func updateOutputWithVariables(variables model.Variables, moduleConfig *config.ModuleConfig) (map[string]bool, error) {
	var (
		variableOpts         = make(map[string]bool)
		encryptVariableNames = sets.NewString()
	)

	for _, s := range variables {
		variableOpts[s.Name] = s.Sensitive

		if s.Sensitive {
			encryptVariableNames.Insert(_variablePrefix + s.Name)
		}
	}

	if encryptVariableNames.Len() == 0 {
		return variableOpts, nil
	}

	reg, err := matchAnyRegex(encryptVariableNames.UnsortedList())
	if err != nil {
		return nil, err
	}

	var shouldEncryptAttr []string

	for k, v := range moduleConfig.Attributes {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		matches := reg.FindAllString(string(b), -1)
		if len(matches) != 0 {
			shouldEncryptAttr = append(shouldEncryptAttr, fmt.Sprintf(`var\.%s`, k))
		}
	}

	// Outputs use encrypted variable should set to sensitive.
	for i, v := range moduleConfig.Outputs {
		if v.Sensitive {
			continue
		}

		reg, err := matchAnyRegex(shouldEncryptAttr)
		if err != nil {
			return nil, err
		}

		if reg.MatchString(string(v.Value)) {
			moduleConfig.Outputs[i].Sensitive = true
		}
	}

	return variableOpts, nil
}

func matchAnyRegex(list []string) (*regexp.Regexp, error) {
	var sb strings.Builder

	sb.WriteString("(")

	for i, v := range list {
		sb.WriteString(v)

		if i < len(list)-1 {
			sb.WriteString("|")
		}
	}

	sb.WriteString(")")

	return regexp.Compile(sb.String())
}
