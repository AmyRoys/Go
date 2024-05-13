package main

import (
	"fmt"
)

func INmain() {
	x := 1
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
