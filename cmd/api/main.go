package main

import (
	"fmt"
	"github.com/johnfercher/go-struct/pkg/domain/entities"
	"github.com/johnfercher/go-struct/pkg/domain/filesystem"
	"github.com/johnfercher/go-struct/pkg/services"
	"log"
)

func main() {
	path := "docs/api"
	projectInterpreter := services.New()
	entities, err := projectInterpreter.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	filtereds := filterEntrypoints(entities)

	for _, filtered := range filtereds {
		fmt.Println(filtered.String())
	}
}

func filterEntrypoints(entityMap map[string]filesystem.Entity) []*entities.Struct {
	var filtered []*entities.Struct
	for _, entity := range entityMap {
		for _, stru := range entity.Structs {
			for _, method := range stru.Methods {
				if method.IsEntrypoint {
					filtered = append(filtered, stru)
					break
				}
			}
		}
	}

	return filtered
}
