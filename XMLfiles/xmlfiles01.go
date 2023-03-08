package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Subject string `xml:"subject"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := os.ReadFile("avengers.xml")

	note := &Notes{}

	xml.Unmarshal([]byte(data), &note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Subject)
	fmt.Println(note.Body)
}