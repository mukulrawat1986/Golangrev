package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Dup prints the count and text of lines that appear more than once
// in the input. It reads from the stdin or from a list of named files.
func Dup(in io.Reader, out io.Writer, fileReader io.ReadWriter) {

	// counts store the string and the count of it
	counts := make(map[string]int)
	var input *bufio.Scanner

	if fileReader == nil {
		// create a scanner object on the in io.Reader
		input = bufio.NewScanner(in)
	} else {
		// create a scanner object on the fileReader
		input = bufio.NewScanner(fileReader)
	}
	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup:%v\n", err)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%s\t%d\n", line, n)
		}
	}
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		Dup(os.Stdin, os.Stdout, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup:%v\n", err)
				continue
			}
			defer f.Close()
			Dup(nil, os.Stdout, f)
		}
	}
}
