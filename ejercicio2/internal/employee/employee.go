package employee

import (
	"ejercicio2/internal/person"
	"fmt"
)

type Employee struct {
	Id       int
	Position string
	person.Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("Id Person: ", e.Person.Id)
	fmt.Println("Name: ", e.Name)
	fmt.Println("DateOfBirth: ", e.DateOfBirth)
	fmt.Println("Id Employee: ", e.Id)
	fmt.Println("Position: ", e.Position)
}
