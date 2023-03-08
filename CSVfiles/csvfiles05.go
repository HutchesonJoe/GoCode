package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"title", "type", "platform"},
		{"Sonic the Hedgehog", "action", "genesis"},
    {"Super Mario", "action", "nes"},
    {"Plants vs Zombies", "tower defense", "pc"},
	}

	f, err := os.Create("videoGames.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)

	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		log.Fatal(err)
	}
}