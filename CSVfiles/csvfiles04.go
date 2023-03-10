package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "role"},
		{"Daniel", "Defoe", "author"},
    {"Neil", "Gaiman", "author"},
    {"Joy", "Harjo", "poet"},
	}

f, err := os.Create("userRecords.csv")
defer f.Close()

if err != nil {
	log.Fatalln("failed to open file", err)
}

w := csv.NewWriter(f)
defer w.Flush()

for _, record := range records {
	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}
}