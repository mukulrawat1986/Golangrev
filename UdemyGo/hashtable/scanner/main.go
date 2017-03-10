package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// an artificial input source
	const input = "It is not the critic who counts; not the man who points out how strong may the doer of deeds could have done them better."

	// create a scanner
	scanner := bufio.NewScanner(strings.NewReader(input))

	// set the split function for the scanning
	scanner.Split(bufio.ScanWords)

	// count the ScanWords
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}
