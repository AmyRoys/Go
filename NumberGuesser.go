package main

import (
    "fmt"
    "math/rand"
)

func NumberGuesser() {
 randomNum := rand.Intn(100)

 for{
	var guess int 
	fmt.Println("Enter your guess: ")
	fmt.Scan(&guess)

	if guess == randomNum{	
		fmt.Println("You guessed it right!")
		break
	}else if guess > randomNum{
		fmt.Println("Your guess is too high")
	}else{
		fmt.Println("Your guess is too low")
	}
 }


}