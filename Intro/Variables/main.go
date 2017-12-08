package main

import "fmt"

// declare global constants
// don't need := here as we are declaring it by having
// const here
const (
	message = "The answer to life is %d\n"
	answer  = 42
)

func main() {
	fmt.Printf(message, answer)
}
