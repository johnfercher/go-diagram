package container

import "github.com/johnfercher/go-diagram/pkg"

type container struct {
	alias    string
	optional pkg.Optional
}

func New(alias string, optional ...pkg.Optional) pkg.Node {
	container := &container{
		alias: alias,
	}

	if len(optional) != 0 {
		container.optional = optional[0]
	}

	return container
}

func (c *container) GetAlias() string {
	return c.alias
}

func (c *container) GetLabel() string {
	return c.optional.Label
}

func (c *container) GetDescription() string {
	return c.optional.Description
}

func (c *container) GetType() string {
	return "Container"
}
