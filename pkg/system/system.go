package system

type system struct {
	alias       string
	label       string
	description string
}

func New(alias string, label string, description string) *system {
	return &system{}
}

func (s *system) GetAlias() string {
	return s.alias
}

func (s *system) GetLabel() string {
	return s.label
}

func (s *system) GetDescription() string {
	return s.description
}
