package main

import (
	"fmt"
	//"time"
	//"math/rand"
	"time"
)

func main() {
	fmt.Println("Starting main func")
	//Declare an unbuffered Channel
	//counter := make(chan int) // exp - 0
	//Declare buffered Channel of capacity 3
	nums := make(chan int)
	v := 0

	go func () {
		// send value in counter
		//sleep := rand.Int63n(1000)
		defer func(){
			fmt.Println("Deferring")
		}()


		//fmt.Println("Starting Ist Go-routine")//
		//v = 1
		//counter <- 10  // exp - 0
		//close(counter) // closing the counter
		fmt.Println("Starting Ist Go-routine")
		fmt.Println("Terminating Ist Go-routine")
	}()

	go func() {
		// send values to buffered channel
		fmt.Println("starting 2nd Go-routine")
		//<- counter  // exp - 0
		nums <- 11

		nums <- 22
		//nums <- 33 // what if nums don`t have 3 values ??
		fmt.Println("terminating 2nd Go-routine")
	}()

	// Read from the un-buffered Channel
	//fmt.Println(<- counter)
	//val, ok := <- counter
	//if ok {
	//	fmt.Println(val)
	//}

	//Read from buffered Channel

	fmt.Println(1)
	fmt.Println(<- nums)
	//time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(v)
	fmt.Println(<- nums)
	//fmt.Println(<- nums) // If nums don`t have 3 values then this will give error
	//time.Sleep(time.Second * 20)  // exp -1
	close(nums)

}