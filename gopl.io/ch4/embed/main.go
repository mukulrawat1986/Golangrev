// example of struct embedding and anonymous fields
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spoke int
}

func main() {

	// use struct literals to initialize
	w := Wheel {
		Circle: Circle {
			Point: Point{
				X: 0,
				Y: 0,
			},
			Radius: 5,
		},
		Spoke : 20,
	}

	fmt.Printf("%#v\n", w)
}
