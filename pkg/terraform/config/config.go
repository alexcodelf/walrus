package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sort"
	"sync"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/seal-io/utils/stringx"
	"github.com/zclconf/go-cty/cty"
	klog "k8s.io/klog/v2"

	transtf "github.com/seal-io/walrus/pkg/templates/translators/terraform"
	"github.com/seal-io/walrus/pkg/terraform/block"
	"github.com/seal-io/walrus/pkg/terraform/convertor"
	"github.com/seal-io/walrus/pkg/terraform/convertors"
)

// Config handles the configuration of resource to terraform config.
type Config struct {
	// File is the hclwrite.File of the Config.
	file *hclwrite.File

	// Attributes is the attributes of the Config.
	// E.g.
	// Attr1 = "xxx"
	// attr2 = 1
	// attr3 = true.
	Attributes map[string]any

	// Blocks blocks like terraform, provider, module, etc.
	/**
	  terraform {
	  	backend "http" {
	  		xxx
	  	}
	  	xxx
	  }
	  provider "aws" {
	  	region = "us-east-1"
	  }

	  module "aws" {
	  	source = "xxx"
	  	region = "us-east-1"
	  }
	*/
	Blocks block.Blocks
}

const (
	// _defaultUsername is the default username in the backend.
	_defaultUsername = "walrus"

	// _updateMethod is the method to update state in the backend.
	_updateMethod = "PUT"
)

// NewConfig returns a new Config.
func NewConfig(ctx context.Context, opts CreateOptions) (*Config, error) {
	// Terraform block.
	var (
		err        error
		attributes map[string]any
	)

	if opts.Attributes != nil {
		attributes = opts.Attributes
	} else {
		attributes = make(map[string]any)
	}

	blocks, err := loadBlocks(ctx, opts)
	if err != nil {
		return nil, err
	}

	c := &Config{
		file:       hclwrite.NewEmptyFile(),
		Attributes: attributes,
		Blocks:     blocks,
	}

	if err = c.validate(); err != nil {
		return nil, err
	}

	// Init the config.
	if err = c.initAttributes(); err != nil {
		return nil, err
	}

	if err = c.initBlocks(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) validate() error {
	for _, b := range c.Blocks {
		if b.Type == "" {
			return fmt.Errorf("invalid b type: %s", b.Type)
		}
	}

	return nil
}

// AddBlocks adds a block to the configuration.
func (c *Config) AddBlocks(blocks block.Blocks) error {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	c.Blocks = append(c.Blocks, blocks...)

	for _, b := range blocks {
		tBlock, err := b.ToHCLBlock()
		if err != nil {
			return err
		}

		c.file.Body().AppendBlock(tBlock)
	}

	return nil
}

// initBlocks initializes the Blocks of the configuration.
func (c *Config) initBlocks() error {
	for i := 0; i < len(c.Blocks); i++ {
		childBlock, err := c.Blocks[i].ToHCLBlock()
		if err != nil {
			return err
		}

		c.file.Body().AppendBlock(childBlock)
		c.file.Body().AppendNewline()
	}

	return nil
}

// initAttributes initializes the attributes of the configuration.
func (c *Config) initAttributes() error {
	if len(c.Attributes) == 0 {
		return nil
	}

	translator := transtf.New()

	attrKeys, attrMap, err := translator.ToOriginalTypeValues(c.Attributes)
	if err != nil {
		return err
	}

	for _, attr := range attrKeys {
		c.file.Body().SetAttributeValue(attr, attrMap[attr])
	}

	return nil
}

// WriteTo writes the config to the writer.
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	// Format the file.
	formatted := hclwrite.Format(Format(c.file.Bytes()))

	return io.Copy(w, bytes.NewReader(formatted))
}

// Reader returns a reader of the config.
func (c *Config) Reader() (io.Reader, error) {
	var buf bytes.Buffer
	if _, err := c.WriteTo(&buf); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}

// Bytes returns the bytes of the config.
func (c *Config) Bytes() ([]byte, error) {
	return hclwrite.Format(Format(c.file.Bytes())), nil
}

// loadBlocks loads the blocks of the configuration.
func loadBlocks(ctx context.Context, opts CreateOptions) (blocks block.Blocks, err error) {
	var (
		tfBlocks       block.Blocks
		providerBlocks block.Blocks
		moduleBlocks   block.Blocks
		variableBlocks block.Blocks
		outputBlocks   block.Blocks
	)
	// Load terraform block.
	if opts.TerraformOptions != nil {
		tfBlocks = block.Blocks{loadTerraformBlock(opts.TerraformOptions)}
	}
	// Other blocks like provider, module, etc.
	// load provider blocks.
	if opts.ProviderOptions != nil {
		providerBlocks, err = loadProviderBlocks(ctx, opts.ProviderOptions)
		if err != nil {
			return nil, err
		}
	}
	// Load module blocks.
	if opts.ModuleOptions != nil {
		moduleBlocks = loadModuleBlocks(opts.ModuleOptions.ModuleConfigs, providerBlocks)
	}
	// Load variable blocks.
	if opts.VariableOptions != nil {
		variableBlocks = loadVariableBlocks(opts.VariableOptions)
	}

	if len(opts.OutputOptions) != 0 {
		outputBlocks = loadOutputBlocks(opts.OutputOptions)
	}

	blocks = make(block.Blocks, 0, block.CountLen(tfBlocks, providerBlocks, moduleBlocks, variableBlocks, outputBlocks))
	blocks = block.AppendBlocks(blocks, tfBlocks, providerBlocks, moduleBlocks, variableBlocks, outputBlocks)

	return blocks, nil
}

// loadTerraformBlock loads the terraform block.
func loadTerraformBlock(opts *TerraformOptions) *block.Block {
	var (
		logger         = klog.Background().WithName("deployer").WithName("tf")
		terraformBlock = &block.Block{
			Type: block.TypeTerraform,
		}
	)

	if opts.ProviderRequirements != nil {
		requiredProviders := &block.Block{
			Type:       block.TypeRequiredProviders,
			Attributes: map[string]any{},
		}
		for provider, requirement := range opts.ProviderRequirements {
			if _, ok := requiredProviders.Attributes[provider]; ok {
				logger.Infof("provider already exists, skip", "provider", provider)
				continue
			}
			pr := make(map[string]any)

			if requirement != nil {
				if len(requirement.VersionConstraints) != 0 {
					pr["version"] = stringx.Join(",", requirement.VersionConstraints...)
				}

				if requirement.Source != "" {
					pr["source"] = requirement.Source
				}
			}
			requiredProviders.Attributes[provider] = pr
		}

		terraformBlock.AppendBlock(requiredProviders)
	}
	backendBlock := &block.Block{
		Type:   block.TypeBackend,
		Labels: []string{"http"},
		Attributes: map[string]any{
			"address": opts.Address,
			// Since the seal server using bearer token and
			// terraform backend only support basic auth.
			// We use the token as the password, and let the username be default.
			"username":               _defaultUsername,
			"skip_cert_verification": opts.SkipTLSVerify,
			// Use PUT method to update the state.
			"update_method":  _updateMethod,
			"retry_max":      10,
			"retry_wait_max": 5,
		},
	}

	terraformBlock.AppendBlock(backendBlock)

	return terraformBlock
}

// loadProviderBlocks returns config providers to get terraform provider config block.
func loadProviderBlocks(ctx context.Context, opts *ProviderOptions) (block.Blocks, error) {
	return convertors.ToProvidersBlocks(ctx, opts.RequiredProviderNames, opts.Connectors, convertor.ConvertOptions{
		SecretMountPath: opts.SecretMonthPath,
		ConnSeparator:   opts.ConnectorSeparator,
		Providers:       opts.RequiredProviderNames,
	})
}

// loadModuleBlocks returns config modules to get terraform module config block.
func loadModuleBlocks(moduleConfigs []*ModuleConfig, providers block.Blocks) block.Blocks {
	var (
		logger       = klog.Background().WithName("deployer").WithName("tf").WithName("config")
		blocks       block.Blocks
		providersMap = make(map[string]any)
	)

	for _, p := range providers {
		alias, ok := p.Attributes["alias"].(string)
		if !ok {
			continue
		}

		if len(p.Labels) == 0 {
			continue
		}
		name := p.Labels[0]
		// Template "{{xxx}}" will be replaced by xxx, the quote will be removed.
		providersMap[name] = fmt.Sprintf("{{%s.%s}}", name, alias)
	}

	for _, mc := range moduleConfigs {
		mb, err := ToModuleBlock(mc)
		if err != nil {
			logger.Infof("get module mb failed, %v", mc)
			continue
		}
		// Inject providers alias to the module.
		if len(mc.SchemaData.Terraform.RequiredProviders) != 0 {
			moduleProviders := map[string]any{}

			for _, p := range mc.SchemaData.Terraform.RequiredProviders {
				if _, ok := providersMap[p.Name]; !ok {
					logger.Infof("provider not found, skip provider: %s", p.Name)
					continue
				}
				moduleProviders[p.Name] = providersMap[p.Name]
			}
			mb.Attributes["providers"] = moduleProviders
		}

		blocks = append(blocks, mb)
	}

	return blocks
}

// loadVariableBlocks returns config variables to get terraform variable config block.
func loadVariableBlocks(opts *VariableOptions) block.Blocks {
	var (
		logger = klog.Background().WithName("terraform").WithName("config")
		blocks = make(block.Blocks, 0, len(opts.Variables)+len(opts.DependencyOutputs))
	)

	// Secret variables.
	for name, sensitive := range opts.Variables {
		blocks = append(blocks, &block.Block{
			Type:   block.TypeVariable,
			Labels: []string{opts.VariablePrefix + name},
			Attributes: map[string]any{
				"type":      "{{string}}",
				"sensitive": sensitive,
			},
		})
	}

	// Dependency variables.
	for k, o := range opts.DependencyOutputs {
		t, err := typeExprTokens(o.Type)
		if err != nil {
			logger.Infof("get type expr tokens failed, %s", err.Error())
			t = hclwrite.TokensForIdentifier("string")
		}

		blocks = append(blocks, &block.Block{
			Type:   block.TypeVariable,
			Labels: []string{opts.ResourcePrefix + k},
			Attributes: map[string]any{
				"type":      tokensToTypeAttr(t),
				"sensitive": o.Sensitive,
			},
		})
	}

	return blocks
}

// loadOutputBlocks returns terraform outputs config block.
func loadOutputBlocks(opts OutputOptions) block.Blocks {
	blockConfig := func(output Output) (string, string) {
		label := fmt.Sprintf("%s_%s", output.ResourceName, output.Name)
		value := fmt.Sprintf(`{{module.%s.%s}}`, output.ResourceName, output.Name)

		return label, value
	}

	// Template output.
	blocks := make(block.Blocks, 0, len(opts))

	for _, o := range opts {
		label, value := blockConfig(o)

		blocks = append(blocks, &block.Block{
			Type:   block.TypeOutput,
			Labels: []string{label},
			Attributes: map[string]any{
				"value":     value,
				"sensitive": o.Sensitive,
			},
		})
	}

	return blocks
}

// ToModuleBlock returns module block for the given module and variables.
func ToModuleBlock(mc *ModuleConfig) (*block.Block, error) {
	var b block.Block

	if mc.Attributes == nil {
		mc.Attributes = make(map[string]any, 0)
	}

	mc.Attributes["source"] = mc.Source
	b = block.Block{
		Type:       block.TypeModule,
		Labels:     []string{mc.Name},
		Attributes: mc.Attributes,
	}

	return &b, nil
}

func CreateConfigToBytes(ctx context.Context, opts CreateOptions) ([]byte, error) {
	conf, err := NewConfig(ctx, opts)
	if err != nil {
		return nil, err
	}

	return conf.Bytes()
}

// tokensToTypeAttr returns the HCL tokens for a type attribute.
func tokensToTypeAttr(tokens hclwrite.Tokens) string {
	return fmt.Sprintf("{{%s}}", tokens.Bytes())
}

// typeExprTokens returns the HCL tokens for a type expression.
func typeExprTokens(ty cty.Type) (hclwrite.Tokens, error) {
	switch ty {
	case cty.String:
		return hclwrite.TokensForIdentifier("string"), nil
	case cty.Bool:
		return hclwrite.TokensForIdentifier("bool"), nil
	case cty.Number:
		return hclwrite.TokensForIdentifier("number"), nil
	case cty.DynamicPseudoType:
		return hclwrite.TokensForIdentifier("any"), nil
	}

	if ty.IsCollectionType() {
		etyTokens, err := typeExprTokens(ty.ElementType())
		if err != nil {
			return nil, err
		}

		switch {
		case ty.IsListType():
			return hclwrite.TokensForFunctionCall("list", etyTokens), nil
		case ty.IsSetType():
			return hclwrite.TokensForFunctionCall("set", etyTokens), nil
		case ty.IsMapType():
			return hclwrite.TokensForFunctionCall("map", etyTokens), nil
		default:
			// Should never happen because the above is exhaustive.
			return nil, fmt.Errorf("unsupported collection type: %s", ty.FriendlyName())
		}
	}

	if ty.IsObjectType() {
		atys := ty.AttributeTypes()
		names := make([]string, 0, len(atys))

		for name := range atys {
			names = append(names, name)
		}

		sort.Strings(names)

		items := make([]hclwrite.ObjectAttrTokens, len(names))

		for i, name := range names {
			value, err := typeExprTokens(atys[name])
			if err != nil {
				return nil, err
			}

			items[i] = hclwrite.ObjectAttrTokens{
				Name:  hclwrite.TokensForIdentifier(name),
				Value: value,
			}
		}

		return hclwrite.TokensForObject(items), nil
	}

	if ty.IsTupleType() {
		etys := ty.TupleElementTypes()
		items := make([]hclwrite.Tokens, len(etys))

		for i, ety := range etys {
			value, err := typeExprTokens(ety)
			if err != nil {
				return nil, err
			}

			items[i] = value
		}

		return hclwrite.TokensForTuple(items), nil
	}

	return nil, fmt.Errorf("unsupported type: %s", ty.GoString())
}
