package main

import "fmt"

func LCmain(){
	ch := make(chan int, 8)
	ch <- 42
	ch <- 7
	<-ch
	ch <- 64
	// number of queued elements = 1 + 1 - 1 + 1 = 2
	fmt.Println(len(ch), cap(ch))
}