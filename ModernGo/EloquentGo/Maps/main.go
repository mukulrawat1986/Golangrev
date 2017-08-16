package main

import "fmt"

func main() {
	// maps in Go are basically hashmaps or dictionaries
	// key => value (key value pair)

	// declaring a map
	x := make(map[string]int)
	x["first"] = 1
	x["second"] = 2
	fmt.Println(x)

	// accessing the key
	fmt.Println(x["first"])

	// accessing a key which does not exist, comma-ok pattern
	if v, ok := x["third"]; ok {
		fmt.Println(`x["third"] is %d\n`, v)
	} else {
		fmt.Println("Key does not exist\n")
	}
}
