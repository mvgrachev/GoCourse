package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Create("/tmp/example.csv")
	checkError("Cannot create file:", err)
	defer file.Close()

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(file)
	w.Comma = ';'
	defer w.Flush()

	for _, record := range records {
		err := w.Write(record)
		checkError("Cannot write to file:", err)
	}
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
