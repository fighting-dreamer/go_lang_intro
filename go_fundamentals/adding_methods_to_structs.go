package main

import (
	"fmt"
)

type Customer struct {
	FirstName, LastName, Email, MobileNum string
}

//Value reciever type Func
func (c Customer) ToString () string {
	return fmt.Sprintf("Customer : FirstName = %s, LastName = %s, Email = %s, MobNum = %s", c.FirstName, c.LastName, c.Email, c.MobileNum)
	//return "Printed a Customer"
}

//pointer reciever type Func
func (c *Customer) setEmail(email string) {
	c.Email = email
}

func main () {
	c := Customer{"Nipun", "Jindal", "wdsgfd@fqes.com", "3214324"}
	fmt.Println(c.ToString())
	c.setEmail("fnfnwnfkwjnkjw@wnfjnw.com")
	fmt.Println(c.ToString())
}