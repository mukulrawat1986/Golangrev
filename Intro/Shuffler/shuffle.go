package shuffler

import "math/rand"

type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(s Shuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}
