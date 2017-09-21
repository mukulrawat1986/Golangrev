package main

import "fmt"

type englishBot struct{}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	// very custom logic for generating spanish greeting
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}
