package main

import "fmt"

type ContactDetails struct {
	email   string
	zipCode string
}

type Person struct {
	firstName      string
	lastName       string
	contactDetails ContactDetails
}

func main() {
	person := Person{
		firstName: "Saroj",
		lastName:  "Bhattarai",
		contactDetails: ContactDetails{
			email:   "medu@edu.com",
			zipCode: "44700",
		},
	}
	personPointer := &person
	personPointer.updateFirstName("Medu")
	person.print()
}

// This is will not work since it is passed by value
// func (p Person) updateFirstName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p *Person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
