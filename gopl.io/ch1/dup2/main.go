/*
* Dup prints the content and text of lines that appear more than once
* in the input. It reads from stdin or a list of named files.
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// create a map to hold the lines and the count of its appearance
	counts := make(map[string]int)

	// files hold the name of files given on the command lines
	files := os.Args[1:]

	// if no input file is given on the command line we will get input from stdin
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// loop over the file names
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue // move to the next file name
			}
			countLines(f, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(r io.Reader, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup: Error while reading %v\n", err)
	}
}
