package main

import (
	"fmt"
)


func main () {
	var x []int // its only declaring but u can`t directly use it, you need to initialize it using "make"
	x = make([]int, 3, 5) // the initial length = 3, capacity = 5
	x = make([]int, 5) // defining the slice and the initial "length" and "capacity" here becomes same as "length"

	fmt.Println(x)
	y := []int{} // same as declaring the slice but not a nil slice `coz it is initialized as empty but initialized atleast
	fmt.Println("y==nil ?  : ", y == nil)
	fmt.Println("length of y : ", len(y))
	if y == nil {
		fmt.Println("Is y is nil right now ? : ", y == nil)
	}

	var yy []int // it is a nil slice
	if yy == nil {
		fmt.Println("yy is nil Slice")
	}

}