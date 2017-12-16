package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}
func main() {
	x, y := 5, 4
	result := Add(x, y)
	fmt.Println("Addition : ", result)

	result = Subtract(x, y)  // Note that no need to use :=, use simple = now coz result defination got completed above.
	fmt.Println("Subtraction : ", result)
}
