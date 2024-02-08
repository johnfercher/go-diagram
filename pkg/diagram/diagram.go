package diagram

import (
	"fmt"
	"github.com/johnfercher/go-diagram/pkg"
)

type diagram struct {
	nodes     map[string]pkg.Node
	relations []*pkg.Relation
}

func New() pkg.Diagram {
	return &diagram{
		nodes: make(map[string]pkg.Node),
	}
}

func (d *diagram) AddNode(node pkg.Node) {
	d.nodes[node.GetAlias()] = node
}

func (d *diagram) AddRel(nodeA pkg.Node, nodeB pkg.Node) {
	d.relations = append(d.relations, &pkg.Relation{
		NodeA: nodeA,
		NodeB: nodeB,
	})
}

func (d *diagram) Generate() (string, error) {
	c4string := begin + "\n"

	for _, node := range d.nodes {
		c4string += fmt.Sprintf(label+"\n", node.GetType(), node.GetAlias(), node.GetLabel(), node.GetDescription())
	}

	c4string += "\n"

	for _, relation := range d.relations {
		c4string += fmt.Sprintf(rel+"\n", relation.NodeA.GetAlias(), relation.NodeB.GetAlias(), "", "")
	}

	c4string += "\n" + end
	return c4string, nil
}
