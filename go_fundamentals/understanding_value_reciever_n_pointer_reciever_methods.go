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


func (c Customer) ChangeEmail(email string){
	c.Email = email // It won	t actually change the email of the customer instance, coz this method recieves the copy of the "Customer" instance (Value Reciever method)
}


func (c *Customer) ChangeEmail_ptr(email string){
	c.Email = email // It won	t actually change the email of the customer instance, coz this method recieves the copy of the "Customer" instance (Value Reciever method)
}

func main(){
	c := Customer{
		FirstName:"Nipun",
		LastName:"Jindal",
		Email:"nipun.jindal@gmail.in",
		MobileNum:"1234567890",
	}
	fmt.Println(c.Email);
	c.ChangeEmail("new.email@gmail.in")
	fmt.Println(c.Email) // No change

	c.ChangeEmail_ptr("new.email@gmail.in")
	fmt.Println(c.Email) // Changed email
}