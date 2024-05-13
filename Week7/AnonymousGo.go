package main

import (
	"fmt"
	"time"
)

func AGmain() {
	go func() {
		fmt.Println("A function with no name and not even a reference")
	}()
	time.Sleep(1 * time.Second)
}
