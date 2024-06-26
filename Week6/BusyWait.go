package main
import (
	"fmt"
	"time"
)
//will print "no channel was ready" until both channels are ready
//select with a default can become a busy wait problem 
func BWmain() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			i++
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			i++
		default:
			fmt.Printf("no channel was ready\n")
		}
	}
}