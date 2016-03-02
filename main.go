package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	f, _ := os.Open("prefs.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		//Print number of records read.
		//  fmt.Println(len(record))

		for value := range record {
			out := exec.Command("git", "--no-pager", "grep", time.Now().String())
			var outbuf, errbuf bytes.Buffer
			out.Stdout = &outbuf
			out.Stderr = &errbuf
			out.Run()
			if err != nil {
				fmt.Printf("%d - %s", value, err.Error())
			}

			if outbuf.String() == "" {
				fmt.Printf("'%s' not found in repository.\n\n", outbuf.String())
			}
			//Print output from Git Grep
			// fmt.Printf(" %s %s - %v\n\n", errbuf.String(), outbuf.String(), record[value])
		}
	}
}
