package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	concepts()
}

func concepts() {
	firstName := "Jim\n\n"
	fmt.Printf("%s", firstName)
	secondName := `Jim\n\n`
	fmt.Printf("%s", secondName)
}
