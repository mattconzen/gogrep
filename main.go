package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("prefs.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		//Print all records
		//  fmt.Println(record)
		//Print number of records read.
		//  fmt.Println(len(record))

		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}
