package main

import (
	"time"
	"fmt"
)

type TeamMember interface {
	PrintName()
	PrintDetails()
}

type Employee struct {
	FirstName, LastName string
	Dob time.Time
	JobTitle, Location string
}

func (e Employee) PrintName() {
	fmt.Println("FirstName : ", e.FirstName, " LastName : ", e.LastName)
}

func (e Employee) PrintDetails () {
	fmt.Printf("FirstName : %s LastName : %s\n", e.FirstName, e.LastName)
	fmt.Printf("Dob : %s\n", e.Dob.String())
	fmt.Printf("Jobtitle : %s, Location : %s", e.JobTitle, e.Location)
	fmt.Println()
}

type Developer struct {
	Employee
	skills []string
}

func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
	fmt.Println("Skills : ")
	for _, v := range(d.skills) {
		fmt.Printf("%s, ", v)
	}
	fmt.Println()
}

type Manager struct {
	Employee
	locations []string
}

func (m Manager) PrintDetails() {
	m.Employee.PrintDetails()
	fmt.Println("Locations being Managed : ")
	for _, v := range(m.locations) {
		fmt.Printf("%s, ", v)
	}
	fmt.Println()
}

type Team struct {
	TemaName string
	Members []TeamMember // NOW you can y=use any of the MAnager or Employee or Developer structs here as a membre of this "slice"
}

func (t Team) PrintDetailsOfTeam() {
	fmt.Println(t.TemaName)
	for _, v := range(t.Members) {
		v.PrintDetails()
	}
}
func main () {

	t := Team {
		"GO-Coders",
		[]TeamMember{
			Developer{
				Employee{
					FirstName:"Nipun",
					LastName:"Jindal",
					Dob:time.Date(1996, 5, 5, 0, 0, 0, 0, time.UTC),
					JobTitle:"SDE-1",
					Location:"Bangalore",
				},
				[]string{
					"go",
					"scala",
					"Java",
					"Python",
					"C-C++",
					"Haskell",
				},
			},
			Manager{
				Employee{
					FirstName:"Nipun",
					LastName:"Jindal",
					Dob:time.Date(1996, 5, 5, 0, 0, 0, 0,time.UTC),
					JobTitle:"SDE-1",
					Location:"Bangalore",
				},
				[]string{
					"bangalore",
					"Delhi",
				},
			},
			Employee{
				FirstName:"Nipun",
				LastName:"Jindal",
				Dob:time.Date(1996, 5, 5, 0, 0, 0, 0, time.UTC),
				JobTitle:"SDE-1",
				Location:"Bangalore",
			},
		},
	}

	t.PrintDetailsOfTeam()
}

