package main

import "fmt"

// interface
type bot interface {
	getGreeting() string
}

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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
