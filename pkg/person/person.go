package person

import "github.com/johnfercher/go-diagram/pkg"

type person struct {
	alias    string
	optional pkg.Optional
}

func New(alias string, optional ...pkg.Optional) pkg.Node {
	person := &person{
		alias: alias,
	}

	if len(optional) != 0 {
		person.optional = optional[0]
	}

	return person
}

func (p *person) GetAlias() string {
	return p.alias
}

func (p *person) GetLabel() string {
	return p.optional.Label
}

func (p *person) GetDescription() string {
	return p.optional.Description
}

func (p *person) GetType() string {
	return "Person"
}
