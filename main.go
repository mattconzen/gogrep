package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	inputFilePtr := flag.String("input", "input.csv",
		`String: A file containing a comma-separated list of
															terms to search for in the repository. (Default: input.csv)`)

	outputFilePtr := flag.String("output", "output.csv",
		`String: File location for the CSV representation of
															 	the results. (Default: output.csv)`)

	repoLocationPtr := flag.String("repo", "", `String: Repository location on disk.`)

	flag.Parse()

	f, _ := os.Open(*inputFilePtr)
	o, _ := os.Create(*outputFilePtr)
	r := csv.NewReader(bufio.NewReader(f))
	w := csv.NewWriter(o)
	repo := string(*repoLocationPtr)

	for {
		record, err := r.Read()

		if err == io.EOF || err != nil {
			break
		}

		//Print number of records read.
		fmt.Println(len(record))
		fmt.Println(repo)

		// Change directory if repo arg has been passed in
		if repo != "" && len(repo) > 0 {
			dir, err := os.Open(repo)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				dir.Chdir()
			}
		}

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
				w.Write([]string{record[value], "0"})
			} else {
				fmt.Printf(" %s %s - %v\n\n", errbuf.String(), outbuf.String(), record[value])
				w.Write([]string{record[value], "1"})
			}
		}

		// Write all buffered data to the output file.
		w.Flush()
	}
}
