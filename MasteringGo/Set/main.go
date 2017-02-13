package main

import "fmt"

// Set is a set created with a map of strings to empty structs
type Set map[string]struct{}

func main() {
	s := make(Set)

	// set the map to an empty struct with empty values
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}

	// get and print items
	fmt.Println(getSetValues(s))

}

func getSetValues(s Set) []string {
	var retval []string
	for k := range s {
		retval = append(retval, k)
	}
	return retval
}
