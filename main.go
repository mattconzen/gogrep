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
    o, _ := os.Create("output.csv")
	r := csv.NewReader(bufio.NewReader(f))
    w := csv.NewWriter(o)
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		//Print number of records read.
		fmt.Println(len(record))

        os.Chdir(string("C:\\Users\\Matthew\\go\\src\\github.com\\mattconzen\\GoExplore"))
        
		for value := range record {
			out := exec.Command("git", "--no-pager", "grep", record[value])
			var outbuf, errbuf bytes.Buffer
			out.Stdout = &outbuf
			out.Stderr = &errbuf
			out.Run()
			if err != nil {
				fmt.Printf("%d - %s", value, err.Error())
			}

			if outbuf.String() == "" {
				fmt.Printf("'%s' not found in repository.\n\n", record[value])
                w.Write( []string{ record[value], "0" } )
			} else {
                fmt.Printf(" %s %s - %v\n\n", errbuf.String(), outbuf.String(), record[value])
                w.Write( []string{ record[value], "1" } )
            }
        }
        // Write all buffered data to the output file.
        w.Flush()
	}
}
