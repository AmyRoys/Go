package main 

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday{
	case today + 0:
		fmt.Println("Today.")
	case today + 1: 
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 5: 
		fmt.Println("In five days.")
	default:
		fmt.Println("Too far away.")
	}

	deferHW()
}

func deferHW(){
	defer fmt.Println("world")
	fmt.Println("hello")
}