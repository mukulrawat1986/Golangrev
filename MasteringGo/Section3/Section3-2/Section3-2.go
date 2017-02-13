// methods can be attached to any type that you create

package main

import "fmt"

type level int

func (lv *level) raiseShieldLevel(i int) {
	if *lv == 0 {
		*lv = 1
	}

	*lv = (*lv) * level(i)
}

func main() {
	sl := new(level)
	sl.raiseShieldLevel(4)
	sl.raiseShieldLevel(5)

	fmt.Println(*sl)
}
