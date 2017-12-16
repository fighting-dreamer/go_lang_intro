package main

import (
	"fmt"
)
type Customer struct{
	FirstName string
	LastName string
	Email string
	MobileNum string
}

type Customer2 struct {
	FirstName, LastName, Email, MobileNum string
} // You can shorten the defination a struct by this

func main () {
	var c Customer
	c.FirstName = "Nipun"
	c.LastName = "Jindal"
	c.Email = "qwerty@abc.com"
	c.MobileNum = "1234567890"

	var mp map[string]string  // declaring a map but not initializing it
	mp = make(map[string]string) // initializing the map
	mp["FirstName"] = c.FirstName
	mp["LastName"] = c.LastName
	mp["Email"] = c.Email
	mp["MobileNum"] = c.MobileNum
	fmt.Println(c)
	fmt.Println(mp)

	c2 := Customer{ FirstName : "Karamvir", LastName : "Jindal", Email : "qwertyasdfg@njknfd.org", MobileNum:"5464758697765"}
	fmt.Println(c2)
	// Notice the comma at the end
	c3 := Customer{ FirstName : "bjabhb", LastName : "Jddqal", Email : "qwertyasdfg@njknfd.org", MobileNum:"5464758697765",}

	c4 := Customer{"billi", "Jindal", "bcwbbcwf@wbfb.edu", "71841749814196433",} // comma is optional but recommended idiomatically

	fmt.Println(c3, c4)

	c5 := Customer{FirstName:"Nipun", MobileNum:"4152552525234"}
	fmt.Println(c5) // it prints only the fields it have, and leave others as "empty strings" or whatever the default value suppoed to be
}



