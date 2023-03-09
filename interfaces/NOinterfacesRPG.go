package main

import (
	"fmt"
)

type Wizard struct{}

func (w Wizard) Defend() string {
	return "Expelliarmus"
}

func (w Wizard) Forget() string {
	return "Obliviate"
}

type Barbarian struct{}

func (b Barbarian) Defend() string {
	return "Dodge"
}

func main() {
	gandalf := Wizard{}
	fmt.Println("Wizard defends: ", gandalf.Defend())

	conan := Barbarian{}
	fmt.Println("Barbarian defends:", conan.Defend())

	fmt.Println("Wizard makes us all forget:", gandalf.Forget())
}