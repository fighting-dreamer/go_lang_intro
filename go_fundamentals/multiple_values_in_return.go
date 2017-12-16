package main

import (
	"fmt"
)

func Swap (x, y string) (xx, yy string) {
	xx = y
	yy = x
	return
}

func main() {
	x, y := "Nipun", "Jindal"
	xx, yy := Swap(x, y)
	fmt.Println("original values : ", x, y)
	fmt.Println("Swapped values : ", xx, yy)
}

