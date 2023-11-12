package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(name string) {
	p.firstName = name
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func main() {
	alex := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jimmy@gmail.com",
			zip:   10,
		},
	}

	name := "A"
	nameRef := &name
	fmt.Println(&nameRef)
	printPointer(nameRef)

	alex.print()

	alex.updateFirstName("Malcolm")

	alex.print()
}
