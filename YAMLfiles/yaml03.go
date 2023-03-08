package main

import (
	"fmt"
     "os"
     "log"

     "gopkg.in/yaml.v3"
)

func main() {
	trees := []string{"elm", "oak", "maple", "sycamore", "chestnut"}

	data, err := yaml.Marshal(&trees)

	if err != nil {
		log.Fatal(err)
	}

	err2 := os.WriteFile("trees.yaml", data, 0744)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("data written")
}