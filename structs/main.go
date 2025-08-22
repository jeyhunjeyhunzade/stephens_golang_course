package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@mail.com",
			zipCode: 94000,
		},
	}

	// without shortcut
	// jimPointer := &jim
	// jimPointer.updateFirstName("Jimmy")

	// with shortcut, still working - if we defing a receiver with pointer (*person) go allows us to reach its pointer
	jim.updateFirstName("Jimmy")

	jim.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// pointerToPerson.firstName = newFirstName // it works as well
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
