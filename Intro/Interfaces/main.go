// Interface
// Interfaces define behavior rather than specific type

package main

import "math/rand"
import "fmt"

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6}
	shuffle(is)
	fmt.Printf("%v\n", is)

	s := stringSlice{"The", "quick", "brown", "fox"}
	shuffle(s)
	fmt.Printf("%v\n", s)
}
