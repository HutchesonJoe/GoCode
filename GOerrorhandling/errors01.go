package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("We don't have the power")
	fmt.Println("Scotty says:", err)
}