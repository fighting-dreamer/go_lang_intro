package main

import (
	//"math/rand"
	"sync"
	"fmt"
	//"time"
)
var wg sync.WaitGroup

type SqValue struct {
	input, value int
}

func randomGenerator(randomNum chan int) {
	defer wg.Done()
	//time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		//randomNum <- rand.Intn(50)
		randomNum <- (i + 1)
	}
	close(randomNum)
}

func computeSquare(randomNum chan int, squaredNum chan SqValue) {
	defer wg.Done()
	var out SqValue
	for num := range randomNum {
		out.input = num
		out.value = (num * num)
		squaredNum <- out
		//squaredNum <- SqValue{input:num, value:(num << 1),} // you can do it like this also
	}
	close(squaredNum) // If you don`t do this then you get problems only `coz the range oven the "squaredNum" channel in infinite loop, kinda gr8
}

func printValues(squaredNum chan SqValue) {
	defer wg.Done()
	for entry := range squaredNum {
		fmt.Printf("input : %d, output : %d\n", entry.input, entry.value)
	}
	//close(squaredNum) // It never get executed `coz the above for loop fo range over "squaredNum" channel never ends
}

func main () {
	wg.Add(3) // `coz otherwise it may happen that your main method just get completed while other haven`t even started
	randomNum := make(chan int) // if you dont do that then when the computeSquare go-routine start while randomGenerator haven`t, problem would be there, `coz range let the forloop running till channel is closed
	squaredNum := make(chan SqValue)

	fmt.Println("Waiting for it to complete")
	go randomGenerator(randomNum)
	go computeSquare(randomNum, squaredNum)
	go printValues(squaredNum)

	wg.Wait()
	fmt.Println("Got Completed, terminating the program!!")
}
