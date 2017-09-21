package main

import "fmt"

func main() {

	// declare a map using var syntax
	// var colors map[string]string

	// declare a map using make
	colors := make(map[string]string)

	// map using literal syntax
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff"

	fmt.Println(colors)
}
