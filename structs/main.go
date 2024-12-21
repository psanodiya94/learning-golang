package main

import (
	"fmt"
)

type contact struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contact   contact
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) update(firstname string, lastname string) {
	(*p).firstname = firstname
	(*p).lastname = lastname
	(*p).contact.email = firstname + "." + lastname + "@live.com"
}

func update(slice []string) {
	slice[1] = "Kevin"
}

func main() {
	john := person{
		firstname: "John",
		lastname:  "Doe",
		contact: contact{
			email:   "john.doe@live.com",
			zipcode: 411057,
		},
	}

	john.print()

	jim := person{
		firstname: "Jim",
		lastname:  "Doe",
		contact: contact{
			email:   "jim.doe@live.com",
			zipcode: 411057,
		},
	}

	jim.print()

	println("Updating John's name to Johnny Doe")

	john.update("Johnny", "Doe")
	john.print()

	// Slices are passed by reference
	slices := []string{"Hello", "World", "!", "How", "are", "you", "?"}

	fmt.Println(slices)
	update(slices)
	fmt.Println(slices)
}
