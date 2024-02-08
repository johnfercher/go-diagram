package system

import "github.com/johnfercher/go-diagram/pkg"

type system struct {
	alias    string
	optional pkg.Optional
}

func New(alias string, optional ...pkg.Optional) pkg.Node {
	person := &system{
		alias: alias,
	}

	if len(optional) != 0 {
		person.optional = optional[0]
	}

	return person
}
func (s *system) GetAlias() string {
	return s.alias
}

func (s *system) GetLabel() string {
	return s.optional.Label
}

func (s *system) GetDescription() string {
	return s.optional.Description
}

func (s *system) GetType() string {
	return "System"
}
