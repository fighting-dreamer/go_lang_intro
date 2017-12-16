package main

import (
	"fmt"
)

//Go don't support inheritence rather favours composition
// so what it means that, you have Embedded types in go-lang

type Customer struct {
	firstname, lastname, email, mob string
}

type Item struct {
	itemname, cost string
}
type Order struct {
	orderId string
	Customer
	Item
	time_it_take_to_reach string
}

func main() {
	o := Order{
		orderId:              "1234567",
		Customer:             Customer{firstname: "Nipun", lastname: "Jindal", email: "email@email.com", mob: "12452364354"},
		Item:                 Item{itemname: "mysore Masala Dosa", cost: "786"},
		time_it_take_to_reach: "1hr",
	}
	fmt.Print(o)
}
