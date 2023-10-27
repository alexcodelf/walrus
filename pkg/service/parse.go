package service

import (
	"context"
	"fmt"
	"reflect"
	"regexp"

	"entgo.io/ent/dialect/sql"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/variable"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/terraform/parser"
)

const (
	// VariablePrefix the prefix of the variable name.
	VariablePrefix = "_walrus_var_"
	// ServicePrefix the prefix of the service output name.
	ServicePrefix = "_walrus_service_"
)

var (
	// _variableReg the regexp to match the variable.
	_variableReg = regexp.MustCompile(`\${var\.([a-zA-Z0-9_-]+)}`)
	// _serviceReg the regexp to match the service output.
	_serviceReg = regexp.MustCompile(`\${service\.([^.}]+)\.([^.}]+)}`)
)

// ParseAttributesOptions options for parse attributes.
type ParseAttributesOptions struct {
	ServiceRevision *model.ServiceRevision

	ServiceName string

	ProjectID     object.ID
	EnvironmentID object.ID
}

// ParseModuleAttributes parse module variables and dependencies.
func ParseModuleAttributes(
	ctx context.Context,
	mc model.ClientSet,
	attributes map[string]any,
	onlyValidated bool,
	opts ParseAttributesOptions,
) (variables model.Variables, outputs map[string]parser.OutputState, err error) {
	var (
		templateVariables        []string
		dependencyServiceOutputs []string
	)

	replaced := !onlyValidated
	templateVariables, dependencyServiceOutputs = parseAttributeReplace(
		attributes,
		templateVariables,
		dependencyServiceOutputs,
		replaced,
	)

	// If service revision has variables that inherit from cloned revision, use them directly.
	if opts.ServiceRevision != nil && len(opts.ServiceRevision.Variables) > 0 {
		for k, v := range opts.ServiceRevision.Variables {
			variables = append(variables, &model.Variable{
				Name:  k,
				Value: crypto.String(v),
			})
		}
	} else {
		variables, err = GetVariables(ctx, mc, templateVariables, opts.ProjectID, opts.EnvironmentID)
		if err != nil {
			return nil, nil, err
		}
	}

	if !onlyValidated {
		outputs, err = GetServiceDependencyOutputsByID(
			ctx,
			mc,
			opts.ServiceRevision.ServiceID,
			dependencyServiceOutputs)
		if err != nil {
			return nil, nil, err
		}

		// Check if all dependency service outputs are found.
		for _, o := range dependencyServiceOutputs {
			if _, ok := outputs[o]; !ok {
				return nil, nil, fmt.Errorf("service %s dependency output %s not found", opts.ServiceName, o)
			}
		}
	}

	return variables, outputs, nil
}

// parseAttributeReplace parses attribute variable ${var.name} replaces it with ${var._variablePrefix+name},
// service reference ${service.name.output} replaces it with ${var._servicePrefix+name}
// and returns variable names and service names.
// Replaced flag indicates whether to replace the module attribute variable with prefix string.
func parseAttributeReplace(
	attributes map[string]any,
	variableNames []string,
	serviceOutputs []string,
	replaced bool,
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
				replaced,
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
			variableRepl := "${var." + VariablePrefix + "${1}}"
			str = _variableReg.ReplaceAllString(str, variableRepl)

			serviceOutputs = append(serviceOutputs, serviceMatched...)
			serviceRepl := "${var." + ServicePrefix + "${1}_${2}}"

			if replaced {
				attributes[key] = _serviceReg.ReplaceAllString(str, serviceRepl)
			}
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
					replaced,
				)
			}
		}
	}

	return variableNames, serviceOutputs
}

// GetVariables get revision used variables from environment and project.
func GetVariables(
	ctx context.Context,
	client model.ClientSet,
	variableNames []string,
	projectID,
	environmentID object.ID,
) (model.Variables, error) {
	var variables model.Variables

	if len(variableNames) == 0 {
		return variables, nil
	}

	nameIn := make([]any, len(variableNames))
	for i, name := range variableNames {
		nameIn[i] = name
	}

	type scanVariable struct {
		Name      string        `json:"name"`
		Value     crypto.String `json:"value"`
		Sensitive bool          `json:"sensitive"`
		Scope     int           `json:"scope"`
	}

	var vars []scanVariable

	err := client.Variables().Query().
		Modify(func(s *sql.Selector) {
			var (
				envPs = sql.And(
					sql.EQ(variable.FieldProjectID, projectID),
					sql.EQ(variable.FieldEnvironmentID, environmentID),
				)

				projPs = sql.And(
					sql.EQ(variable.FieldProjectID, projectID),
					sql.IsNull(variable.FieldEnvironmentID),
				)

				globalPs = sql.IsNull(variable.FieldProjectID)
			)

			s.Where(
				sql.And(
					sql.In(variable.FieldName, nameIn...),
					sql.Or(
						envPs,
						projPs,
						globalPs,
					),
				),
			).SelectExpr(
				sql.Expr("CASE "+
					"WHEN project_id IS NOT NULL AND environment_id IS NOT NULL THEN 3 "+
					"WHEN project_id IS NOT NULL AND environment_id IS NULL THEN 2 "+
					"ELSE 1 "+
					"END AS scope"),
			).AppendSelect(
				variable.FieldName,
				variable.FieldValue,
				variable.FieldSensitive,
			)
		}).
		Scan(ctx, &vars)
	if err != nil {
		return nil, err
	}

	found := make(map[string]scanVariable)
	for _, v := range vars {
		ev, ok := found[v.Name]
		if !ok {
			found[v.Name] = v
			continue
		}

		if v.Scope > ev.Scope {
			found[v.Name] = v
		}
	}

	// Validate module variable are all exist.
	foundSet := sets.NewString()
	for n, e := range found {
		foundSet.Insert(n)
		variables = append(variables, &model.Variable{
			Name:      n,
			Value:     e.Value,
			Sensitive: e.Sensitive,
		})
	}
	requiredSet := sets.NewString(variableNames...)

	missingSet := requiredSet.
		Difference(foundSet).
		Difference(WalrusMetadataSet)
	if missingSet.Len() > 0 {
		return nil, fmt.Errorf("missing variables: %s", missingSet.List())
	}

	return variables, nil
}
