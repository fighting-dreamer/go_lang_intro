package main

import (
	"fmt"
)

func SplitValues(f func (int) (int, int), x int) (int, int){
	return f(x)
}

func main () {
	//fmt.Println("")
	x := 50
	y := 10
	f := func (sum int) (int, int) {
		return sum * y / 2, sum * 2 / y;
	}

	a, b := SplitValues(f, x)
	fmt.Println(a, b)
}
