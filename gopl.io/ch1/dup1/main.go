// Dup prints the text of each line that appears more than once in
// the standard input, preceeded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map to keep a count of the lines and the number of
	// times the line is repeated.
	counts := make(map[string]int)

	// create a Scanner object on the standard input
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup: Error while scanning %v", err)
		os.Exit(1)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
