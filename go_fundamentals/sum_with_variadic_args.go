package main


import (
	"fmt"
)

func Sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main () {
	fmt.Println("Sum with variadic Args with Array : ")
	total := Sum(2,3,4,5,1,6,0)
	fmt.Println("Sum of 2,3,4,5,1,6,0 = ", total)

	fmt.Println("Sum with variadic args with Slice : ")
	slice_v := []int{2,3,4,5,6,7}
	total = Sum(slice_v...)
	fmt.Println("Sum for 2,3,4,5,6,7 = ", total)
}