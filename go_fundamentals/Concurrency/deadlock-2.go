package main

import "fmt"

func f1(counter, num chan int) {
	counter <- 10
	fmt.Printf("F1 : %d", <-num)
}

func main() {
	var num, counter chan int
	f1(counter, num)
	fmt.Printf("Main : %d", <-counter)
	num <- 20
	fmt.Println("deadlock it seems")
}
