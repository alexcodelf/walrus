package resourcerun

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	pkgresource "github.com/seal-io/walrus/pkg/resource"
	"github.com/seal-io/walrus/pkg/templates/openapi"
	"github.com/seal-io/walrus/pkg/templates/translator"
	"github.com/seal-io/walrus/pkg/terraform/config"
)

const ResourceRunIDLabel = "walrus.seal.io/resource-run-id"

func getModuleConfig(
	run *model.ResourceRun,
	template *model.TemplateVersion,
	opts *ConfigLoaderOptions,
) (*config.ModuleConfig, error) {
	mc := &config.ModuleConfig{
		Name:   opts.Context.Resource.Name,
		Source: template.Source,
	}

	if template.Schema.IsEmpty() {
		return mc, nil
	}

	mc.SchemaData = template.Schema.TemplateVersionSchemaData

	if template.Schema.OpenAPISchema == nil ||
		template.Schema.OpenAPISchema.Components == nil ||
		template.Schema.OpenAPISchema.Components.Schemas == nil {
		return mc, nil
	}

	// Variables.
	var (
		variablesSchema    = template.Schema.VariableSchema()
		outputsSchemas     = template.Schema.OutputSchema()
		sensitiveVariables = sets.Set[string]{}
	)

	if variablesSchema != nil {
		attrs, err := translator.ToGoTypeValues(run.Attributes, *variablesSchema)
		if err != nil {
			return nil, err
		}

		mc.Attributes = attrs

		for n, v := range variablesSchema.Properties {
			// Add sensitive from schema variable.
			if v.Value.WriteOnly {
				sensitiveVariables.Insert(fmt.Sprintf(`var\.%s`, n))
			}

			if n == openapi.WalrusContextVariableName {
				mc.Attributes[n] = opts.Context
			}
		}
	}

	// Outputs.
	if outputsSchemas != nil {
		sps := outputsSchemas.Properties
		mc.Outputs = make([]config.Output, 0, len(sps))

		sensitiveVariableRegex, err := matchAnyRegex(sensitiveVariables.UnsortedList())
		if err != nil {
			return nil, err
		}

		for k, v := range sps {
			origin := openapi.GetExtOriginal(v.Value.Extensions)
			co := config.Output{
				Sensitive:    v.Value.WriteOnly,
				Name:         k,
				ResourceName: opts.Context.Resource.Name,
				Value:        origin.ValueExpression,
			}

			if !v.Value.WriteOnly {
				// Update sensitive while output is from sensitive data, like secret.
				if sensitiveVariables.Len() != 0 && sensitiveVariableRegex.Match(origin.ValueExpression) {
					co.Sensitive = true
				}
			}

			mc.Outputs = append(mc.Outputs, co)
		}
	}

	return mc, nil
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

// getSubject gets the subject of the given resource.
func getSubject(ctx context.Context, mc model.ClientSet, res *model.Resource) (*model.Subject, error) {
	var (
		subjectID object.ID
		err       error
	)

	s, _ := session.GetSubject(ctx)
	if s.ID != "" {
		subjectID = s.ID
	} else {
		subjectID, err = pkgresource.GetSubjectID(res)
		if err != nil {
			return nil, err
		}
	}

	if subjectID == "" {
		return nil, fmt.Errorf("subject id is empty")
	}

	return mc.Subjects().Get(ctx, subjectID)
}
