package main

import (
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

	for _, entity := range entities {
		entity.Print("")
	}
}
