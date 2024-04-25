package api

import (
	"github.com/seal-io/walrus/pkg/templates/openapi"
)

const WalrusContextVariableName = openapi.WalrusContextVariableName

// Context indicates the walrus-related metadata,
// will set to attribute context while user module include this attribute.
type Context struct {
	// Project indicates the project metadata.
	Project struct {
		Name string `json:"name,omitempty"`
	} `json:"project,omitempty"`

	// Environment indicate the environment metadata.
	Environment struct {
		Name string `json:"name,omitempty"`
	} `json:"environment,omitempty"`

	// Resource indicates the resource metadata.
	Resource struct {
		Name string `json:"name,omitempty"`
	} `json:"resource,omitempty"`
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetProject(name string) *Context {
	c.Project.Name = name

	return c
}

func (c *Context) SetEnvironment(name string) *Context {
	c.Environment.Name = name

	return c
}

func (c *Context) SetResource(name string) *Context {
	c.Resource.Name = name

	return c
}
