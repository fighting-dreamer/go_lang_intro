package main

import (
	"fmt"
	"time"
)

type TeamMember interface{
	PrintName()
	PrintDetails()
}

type Employee struct {
	FirstName, LastName string
	Dob time.Time
	JobTitle, Location string
}

func (e Employee) PrintName() {
	fmt.Printf("Employee Name : %s %s", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Employee Name : %s %s\n", e.FirstName, e.LastName)
	fmt.Printf("Dob : %s\n", e.Dob)
	fmt.Printf("JobTitle : %s\nLocation : %s\n", e.JobTitle, e.Location)
}

type Developer struct {
	FirstName string
	Employee
	Skills []string
}

//Overriding the Employee`s method
func (d Developer) PrintDetails() {
	//Calling the Employee`s PrintDetails
	d.Employee.PrintDetails()
	fmt.Print("Developer Skills : ")
	for _, v := range(d.Skills) {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}

func main () {
	d := Developer{FirstName:"Nipun",
		           Employee : Employee{FirstName:"Varun", LastName:"Jindal", Dob:time.Now(), JobTitle:"SDE-1", Location:"Bangalore"},
		           Skills:[]string{"GO-lang", "C-C++", "Java", "python", "scala"}}
	d.PrintDetails() // You are able to call the Employee`s method from Developer Object, note the result hae FirstName diff from what we had for Developer
	fmt.Println(d.FirstName)
	fmt.Println(d.Employee.FirstName)
	d.PrintName() // It directly picked up the Employee`s PrintName and internally called "d.Employee.PrintName()"

}
