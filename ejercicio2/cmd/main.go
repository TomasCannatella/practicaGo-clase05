package main

import (
	"ejercicio2/internal/employee"
	"ejercicio2/internal/person"
)

func main() {
	person := person.Person{
		Id:          1,
		Name:        "Tomas Cannatella",
		DateOfBirth: "22/06/2002",
	}

	employee := employee.Employee{
		Id:       113069,
		Position: "Software Developer",
		Person:   person,
	}

	employee.PrintEmployee()
}
