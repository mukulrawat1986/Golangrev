// dup2 prints the count and test of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		Count(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			Count(f, counts)
		}
	}
}

func Count(r io.Reader, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func Print(w io.Writer, counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(w, "%d\t%s\n", n, line)
		}
	}
}
