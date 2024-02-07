package main

import (
	"fmt"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/johnfercher/go-struct/pkg/domain/entities"
	"github.com/johnfercher/go-struct/pkg/domain/filesystem"
	"github.com/johnfercher/go-struct/pkg/services"
	"github.com/johnfercher/go-tree/node"
	"log"
)

type Element struct {
	Label string
	Type  string
}

func main() {
	path := "docs/api"
	projectInterpreter := services.New()
	projectEntities, err := projectInterpreter.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	filtereds := filterEntrypointMethods(projectEntities)

	var entrypoints []*diagram.Node
	for _, filtered := range filtereds {
		entrypoints = append(entrypoints, gcp.Compute.ComputeEngine(diagram.NodeLabel(fmt.Sprintf("%s_%s", filtered.Struct, filtered.Name))))
	}

	d, err := diagram.New(diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(entrypoints)-1; i++ {
		d.Connect(entrypoints[i], entrypoints[i+1], diagram.Forward())
	}

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

func buildTree() *node.Node[Element] {

}

func filterEntrypointMethods(entityMap map[string]filesystem.Entity) []*entities.Function {
	var filtered []*entities.Function
	for _, entity := range entityMap {
		for _, stru := range entity.Structs {
			for _, method := range stru.Methods {
				if method.IsEntrypoint {
					filtered = append(filtered, method)
					break
				}
			}
		}
	}

	return filtered
}
