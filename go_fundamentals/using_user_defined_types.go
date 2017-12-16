package main

import (
	"fmt"
)

type Address struct {
	street, district, city string
	zip                    int64
}

type Order struct {
	firstname, lastname string
	phone               int64
	address             []Address
}

func main() {
	o := Order{
		firstname: "Nipun",
		lastname:  "Jindal",
		phone:     1123454354,
		address: []Address{
			Address{
				street:   "qwerty",
				district: "btm",
				city:     "Bangalore",
				zip:      12345678,
			},
		},
	}

	fmt.Println(o)
}
