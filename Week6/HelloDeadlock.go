package main

import "fmt"
//results in deadlock because the main function is waiting for a message to be sent to the channel
func HDmain(){
	messages := make(chan string)

	<-messages
	go func(){
		messages <-"Hello"}()
		fmt.Println(<-messages)
}