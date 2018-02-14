package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[:19])
}
