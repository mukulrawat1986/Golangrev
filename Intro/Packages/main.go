package main

import (
	"fmt"

	"github.com/mukulrawat1986/Golangrev/Intro/Shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type weightString struct {
	weight int
	s      string
}

type stringSlice []weightString

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringSlice) Weight(i int) int {
	return s[i].weight
}

func main() {
	i := intSlice{3, 1, 4, 1, 5, 9}
	shuffler.Shuffle(i)
	fmt.Printf("%v\n", i)

	i1 := stringSlice{
		weightString{100, "Hello"},
		weightString{200, "World"},
		weightString{10, "Goodbye"},
	}

	shuffler.WeightedShuffle(i1)
	fmt.Printf("%v\n", i1)

}
