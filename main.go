package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"flag"
)

func main() {
	inputFilePtr := flag.String("input", "input.csv",
															`String: A file containing a comma-separated list of
															terms to search for in the repository. (Default: input.csv)`)

  outputFilePtr := flag.String("output", "output.csv",
															 `String: File location for the CSV representation of
															 	the results. (Default: output.csv)`)

  repoLocationPtr := flag.String("repo", "C:\repo", `String: Repository location on disk.`)

	flag.Parse()

	f, _ := os.Open(*inputFilePtr)
  o, _ := os.Create(*outputFilePtr)
	r := csv.NewReader(bufio.NewReader(f))
  w := csv.NewWriter(o)

	for {
		record, err := r.Read()

		if err == io.EOF || err != nil {
			break
		}

		//Print number of records read.
		fmt.Println(len(record))

    os.Chdir(string(*repoLocationPtr))

		for value := range record {
			out := exec.Command("git", "--no-pager", "grep", "-i", record[value])
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
