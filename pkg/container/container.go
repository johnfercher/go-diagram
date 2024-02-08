package container

type container struct {
	alias       string
	label       string
	description string
}

func New(alias string, label string, description string) *container {
	return &container{
		alias:       alias,
		label:       label,
		description: description,
	}
}

func (c *container) GetAlias() string {
	return c.alias
}

func (c *container) GetLabel() string {
	return c.label
}

func (c *container) GetDescription() string {
	return c.description
}
