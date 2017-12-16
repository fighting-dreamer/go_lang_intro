package main

import (
	"fmt"
)

func main () {
	x := make(map[int]int)
	x[5] = 10
	fmt.Println(x[5])

	y := map[int]string{}
	y[0] = "I m getting maps in go lang"
	y[1] = "hey it is cool"
	y[5] = "It is actually working"
	fmt.Println(y, len(y))
	fmt.Println(y[0])

	// working with maps :
	value, ok := y[0]
	if ok {
		fmt.Println(value)
	}

	if value, ok = y[5]; ok {
		fmt.Println(value, ok)
	}

	value, ok = y[3] // 'coz y[3] don`t exist, it gives value as empty string
	fmt.Println(value, ok , "\nIs value nil ? :", value == "")

//	deleting some values in the maps

	//deleting y[5]
	delete(y, 5) // you mention the map_identifier and the key for which you want to delete
	// let`s see if we can delete more than on key-value pairs in maps at once ? :
	fmt.Println("map 'y' before multiple deletes : ", y)
	y[2] = "qwerty"
	y[3] = "qwertyu"
	y[4] = "chbfkhwbhvkwbvw"
	//delete(y, 2, 3) // actually you can delete one elem at a time in maps so
	delete(y, 4)
	fmt.Println("map 'y' after multiple deletes : ", y)

	//iterating over maps
	for key, val := range y {  // note while iterating, the key comes first then value so
		fmt.Printf("y[%d] : %s\n", key, val)
	} // order in which map is getting iterated will not be same every time

}
