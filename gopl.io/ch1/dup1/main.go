// Dup1 prints the text of each line that appears more than once
// in the standard input, preceeded by its count.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// create a map to hold the count
	counts := make(map[string]int)

	// call the function to count
	Count(counts, os.Stdin)

	// call the function to print the count
	Print(os.Stdout, counts)
}

func Count(counts map[string]int, r io.Reader) {
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
