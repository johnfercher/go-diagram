package pkg

type Diagram interface {
	AddNode(node Node)
	AddRel(nodeA Node, nodeB Node)
	Generate() (string, error)
}

type Node interface {
	GetAlias() string
	GetLabel() string
	GetDescription() string
	GetType() string
}

type Relation struct {
	NodeA    Node
	NodeB    Node
	Optional Optional
}

type Optional struct {
	Label       string
	Description string
}
