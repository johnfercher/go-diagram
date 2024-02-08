package rel

import "github.com/johnfercher/go-diagram/pkg"

type rel struct {
	nodeA pkg.Node
	nodeB pkg.Node
	label string
	tech  string
}

func New(nodeA pkg.Node, nodeB pkg.Node, label string, tech string) *rel {
	return &rel{
		nodeA: nodeA,
		nodeB: nodeB,
		label: label,
		tech:  tech,
	}
}
