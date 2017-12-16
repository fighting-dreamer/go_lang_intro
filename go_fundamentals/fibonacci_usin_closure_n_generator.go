package main

import (
	"fmt"
)



func main() {
	fmt.Println("")
	x := 0
	y := 1
	fibonacci_generator := func () (result int) {
		result = x + y
		x = y
		y = result
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci_generator())
	}
}

