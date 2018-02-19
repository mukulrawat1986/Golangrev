package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := make([]string, 0, 4)
	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	printer(words)
}
