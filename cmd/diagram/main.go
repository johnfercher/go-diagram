package main

import (
	"fmt"
	"github.com/johnfercher/go-diagram/pkg/container"
	"github.com/johnfercher/go-diagram/pkg/diagram"
	"github.com/johnfercher/go-diagram/pkg/person"
	"github.com/johnfercher/go-diagram/pkg/system"
	"log"
)

func main() {
	d := diagram.New()

	system1 := system.New("system1")
	system2 := system.New("system2")
	system3 := system.New("system3")
	container1 := container.New("container1")
	container2 := container.New("container2")
	container3 := container.New("container3")
	person := person.New("person")

	d.AddNode(system1)
	d.AddNode(system2)
	d.AddNode(system3)
	d.AddNode(container1)
	d.AddNode(container2)
	d.AddNode(container3)
	d.AddNode(person)

	d.AddRel(person, system1)
	d.AddRel(system1, system2)
	d.AddRel(system1, system3)
	d.AddRel(system2, container1)
	d.AddRel(system2, container2)
	d.AddRel(system3, container3)

	content, err := d.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
