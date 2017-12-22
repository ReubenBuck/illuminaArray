package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := 8; i < len(records); i++ {
		if records[i][0] == "[Controls]" {
			break
		}
		fmt.Print(">", records[i][0], "\n")
		fmt.Print(records[i][5], "\n", "\n")
	}

}
