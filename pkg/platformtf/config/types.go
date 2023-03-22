package config

import (
	"github.com/seal-io/seal/pkg/dao/model"
)

const (
	// BlockTypeTerraform represents the terraform block.
	BlockTypeTerraform = "terraform"
	// BlockTypeBackend represents the backend block inside terraform block.
	BlockTypeBackend = "backend"
	// BlockTypeRequiredProviders represents the required_providers block inside terraform block.
	BlockTypeRequiredProviders = "required_providers"

	// BlockTypeProvider represents the provider block.
	BlockTypeProvider = "provider"
	// BlockTypeModule represents the module block.
	BlockTypeModule = "module"
	// BlockTypeVariable represents the variable block.
	BlockTypeVariable = "variable"
	// BlockTypeOutput represents the output block.
	BlockTypeOutput = "output"
	// BlockTypeResource represents the resource block.
	BlockTypeResource = "resource"

	// BlockTypeK8s represents the kubernetes block.
	BlockTypeK8s = "kubernetes"
)

const (
	ProviderK8s  = "kubernetes"
	ProviderHelm = "helm"
)

const (
	FileMain = "main.tf"
	FileVars = "terraform.tfvars"
)

// ModuleConfig is a struct with model.Module and its variables.
type ModuleConfig struct {
	// Name is the name of the app module relationship.
	Name          string
	ModuleVersion *model.ModuleVersion
	// Attributes is the attributes of the module.
	Attributes map[string]interface{}
}

// CreateOptions represents the CreateOptions of the Config.
type CreateOptions struct {
	Attributes       map[string]interface{}
	TerraformOptions *TerraformOptions
	ProviderOptions  *ProviderOptions
	ModuleOptions    *ModuleOptions
	VariableOptions  *VariableOptions
}

type (
	// TerraformOptions is the options to create terraform block.
	TerraformOptions struct {
		// Token is the backend token to authenticate with the Seal Server of the terraform config.
		Token string
		// Address is the backend address of the terraform config.
		Address string
		// SkipTLSVerify is the backend cert verification of the terraform config.
		SkipTLSVerify bool
	}

	// ProviderOptions is the options to create provider blocks.
	ProviderOptions struct {
		// SecretMountPath is the mount path of the secret in the terraform config.
		SecretMonthPath string
		// ConnectorSeparator is the separator of the terraform provider alias name and id.
		ConnectorSeparator string
		// RequiredProviders is the required providers of the terraform config.
		// e.g. ["kubernetes", "helm"]
		RequiredProviders []string
		Connectors        model.Connectors
	}

	// ModuleOptions is the options to create module blocks.
	ModuleOptions struct {
		// ModuleConfigs is the module configs of the deployment.
		ModuleConfigs []*ModuleConfig
	}

	// VariableOptions is the options to create variables blocks.
	VariableOptions struct {
		// Secrets is the  model.Secrets of the deployment.
		Secrets model.Secrets
		// Variables is the variables of the deployment.
		Variables map[string]interface{}
	}
)

type ProviderConvertOptions struct {
	SecretMountPath string
	ConnSeparator   string
}
