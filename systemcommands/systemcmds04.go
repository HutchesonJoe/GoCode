package main

import (
    "fmt"
    "io"
    "log"
    "os/exec"
    "strings"
)

func upper(data string) string {
	return strings.ToUpper(data)
}

func main() {
	cmd := exec.Command("echo", "the third moon of endor")

	stdout, err := cmd.StdoutPipe()
  fmt.Println(stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", upper(string(data)))
}