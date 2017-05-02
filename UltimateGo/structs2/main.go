package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type lisa struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Declare a variable of anonymous type and init using a
	// struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	var b bill
	var l lisa

	// explicit conversion
	b = bill(l)

	fmt.Println(b, l)

	// Display the values.
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
