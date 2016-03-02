package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
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
			out := exec.Command("git", "--no-pager grep 'range'")
			var outbuf, errbuf bytes.Buffer
			out.Stdout = &outbuf
			out.Stdin = &errbuf
			out.Run()
			if err != nil {
				fmt.Printf(err.Error())
			}
			fmt.Printf(" %s %s - %v\n", errbuf.String(), outbuf.String(), record[value])
		}
	}
}
