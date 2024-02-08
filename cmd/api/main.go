package main

import (
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

	var entrypoints *node.Node[Element]
	for _, filtered := range filtereds {
		if filtered.Struct != "" {

		}
		entrypoints = append(entrypoints, entrypoint)
	}

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

func buildTree() *node.Node[Element] {
	return nil
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
