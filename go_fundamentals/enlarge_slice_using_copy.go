package main

import (
	"fmt"
)


func main () {
	x := []int{2,3,4}
	fmt.Println("Slice x : ", x, "length ", len(x), "Capacity ", cap(x))

	// new Sclice
	y := make([]int, 5, 10) // y is slice of length 5 and capacity = 10
	fmt.Println("Slice y (original): ", y, "length ", len(y), "Capacity ", cap(y))

	// copying x to y
	copy(y, x) // copy "x" in "y"
	fmt.Println("Slice y (after copy) : ", y, "length ", len(y), "Capacity ", cap(y))
	//set the left values of y
	y[3] = 20
	y[4] = 50
	fmt.Println("Slice y (filling missing values): ", y, "length ", len(y), "Capacity ", cap(y))
}