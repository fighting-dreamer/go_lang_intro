package main

import (
	"fmt"
)

func main () {
	// Declaring a channel
	counter := make(chan int)
	//sending 10 as input in the "counter" channel
	go func () {
		counter <- 10
	}()

	fmt.Println(<- counter) // recieving the "counter" channel value
	// this will solve the deadlock we had in earlier program(deadlock_using_un-buffered_channles)
}
