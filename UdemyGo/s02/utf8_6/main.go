package main

import "fmt"

func main() {

	var runes []rune

	for _, r := range "bye こんにちは世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	fmt.Printf("%q\n", []rune("bye こんにちは世界"))
}
