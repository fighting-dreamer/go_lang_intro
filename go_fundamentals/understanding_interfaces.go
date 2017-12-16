package main

import (
	"fmt"
	"time"
)

type TeamMember interface { // defineing Interface
	PrintName()
	PrintDetails()
}

type Employee struct {
	Firstname, Lastname string
	Dob time.Time
	JobTitle, Location string
}

func (e Employee) PrintName (){
	fmt.Printf("Employee Name: %s %s\n", e.Firstname, e.Lastname)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Employee Name : %s %s\n", e.Firstname, e.Lastname)
	fmt.Printf("Date of Birth : %s\n", e.Dob.String()) // cool part
	fmt.Printf("Job Title : %s \nLocation : %s\n", e.JobTitle, e.Location)
}
func main () {
	e1 := Employee{"Nipun", "Jindal", time.Date(1990, time.February, 17, 0, 0, 0, 0, time.UTC), "SDE-1", "Bangalore"}
	e1.PrintName()
	e1.PrintDetails()

	fmt.Println("\n")

	// using current time
	e2 := Employee{"Nipun", "Jindal", time.Now(), "SDE-1", "Bangalore"}
	e2.PrintName()
	e2.PrintDetails()
}
