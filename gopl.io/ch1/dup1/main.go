package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Dup prrints the text of each line that appears more than once
// in the standard input, preceeded by its count.
func Dup(in io.Reader, out io.Writer) {
	counts := make(map[string]int)

	// create a bufio scanner over the input reader
	input := bufio.NewScanner(in)
	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup1:%v\n", err)
		os.Exit(1)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(out, "%s\t%d\n", line, n)
		}
	}
}

func main() {

}
