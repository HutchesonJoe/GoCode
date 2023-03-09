package main

import (
	"fmt"
	"reflect"
)

type Wizard struct{}

func (w Wizard) Defend() string {
	return "Expelliarmus"
}

func (w Wizard) Forget() string {
	return "Obliviate"
}

func (w Wizard) Attack() string {
	return "Bam"
}

type Barbarian struct{}

func (b Barbarian) Defend() string {
	return "Dodge"
}

func (b Barbarian) Attack() string {
	return "Bam"
}

type Player interface {
	Defend() string
	Attack() string
}

func main() {
	gandalf := Wizard{}

	conan := Barbarian{}
	
	players := []Player{gandalf, conan}

	for _, a := range players {
		fmt.Println(reflect.TypeOf(a).Name(), "defends:", a.Defend())
		fmt.Println(reflect.TypeOf(a).Name(), "attacks:", a.Attack())
	}

	fmt.Println("Wizard makes us all forget:", gandalf.Forget())
}