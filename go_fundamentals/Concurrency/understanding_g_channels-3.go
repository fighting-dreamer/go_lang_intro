package main

import (
	"fmt"
	//"time"
)

// testing what happen If I send and recieve on a closed channel



func main () {
	counter := make(chan int)
	go func() {
		//time.Sleep(time.Second)// If you do this, the program in main will halt at the unbuffered-channel`s(counter `s) receive, coz it can`t get executed till then
		counter <- 10
		close(counter)
	}()

	// Reading from a closed un-buffered channel
	fmt.Println(counter) // printing the channel
	fmt.Println(<- counter) // actual receiving from channel
	val, ok := <- counter // trying to receive again from it

	if ok {
		fmt.Println(val) // this won`t execute, `coz "ok" is false
	}
}
