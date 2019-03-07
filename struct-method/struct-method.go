package main

import (
	"fmt"
)

type person struct {
	Name string
}

type android struct {
	person
	Model string
}

func (p *person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	//p := new(person)
	//p.Name = "Rian"
	p := person{"Rian"}
	p.talk()

	a := new(android)
	a.person.Name = "Lili"
	a.person.talk()
}
