package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Dup prints the count and text of lines that appear more than once
// in the input. It reads from the stdin or from a list of named files.
func Dup(in io.Reader, out io.Writer) {

	// counts store the string and the count of it
	counts := make(map[string]int)

	// create a scanner object on the in io.Reader
	input := bufio.NewScanner(in)

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

}
