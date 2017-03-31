// updating a variable passed to function indirectly

package main

import "fmt"

func f(p *int) int {
	*p++
	return *p
}

func main() {
	v := 2
	f(&v)
	fmt.Println(f(&v))

}
