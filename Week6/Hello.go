package main

import "fmt"

func Hmain(){
	messages := make(chan string)
	//Deadlocks here because the main function is waiting for a message to be sent to the channel
	go func(){
		messages <-"Hello1"}()
	//This line is executed before the message is sent to the channel
	go func(){
		messages <-"Hello2"}()
	fmt.Println(<-messages)
}