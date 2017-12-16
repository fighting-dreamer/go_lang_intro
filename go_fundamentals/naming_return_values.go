package main

import (
	"fmt"
)

func Add_with_naming_values(x, y int) (result int) {
	result = x + y
	return
}

func main () {
	x,y := 5,4
	result := Add_with_naming_values(x,y)
	fmt.Println("Addition : ", result)
}

