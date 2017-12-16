package main

import (
	"fmt"
)

func InfinoteLoop () {
	for {fmt.Println("nothing")}
}

func TestingDefer() {
	var x int = 5
	defer fmt.Println("deferred call -- 1 : X = ", x) // note that In this "x" will be "5" when it be called
	x = 10
	defer fmt.Println("deferred call -- 2 : X = ", x)
	defer InfinoteLoop() // It paniced only then it started the infinite loop thus it just store expression and required values, not the value of expression
	panic (x)
	fmt.Println("It actually panicked!!")
}

func main () {
	TestingDefer()
}
