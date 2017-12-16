package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printCounts(sqNum <-chan int, doubleNum <-chan int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		select {
			case s := <-sqNum :
				fmt.Printf("sqNum %d\n", s)
			case d := <-doubleNum :
				fmt.Printf("Double Num : %d\n", d)
		}
	}
}

func generateSqNum(sqNum chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		sqNum <- (i + 1)
	}
}

func generateDouble(doubleNum chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		doubleNum <- (i + 1) * 2
	}
}

func main() {
	wg.Add(3)
	sqNum := make(chan int)
	doubleNum := make(chan int)

	go generateDouble(doubleNum)
	go generateSqNum(sqNum)

	go printCounts(sqNum, doubleNum)

	wg.Wait()
	fmt.Println("terminating the program")
}