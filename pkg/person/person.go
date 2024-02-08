package person

type person struct {
	alias       string
	label       string
	description string
}

func New(alias string, label string, description string) *person {
	return &person{
		alias:       alias,
		label:       label,
		description: description,
	}
}

func (p *person) GetAlias() string {
	return p.alias
}

func (p *person) GetLabel() string {
	return p.label
}

func (p *person) GetDescription() string {
	return p.description
}
