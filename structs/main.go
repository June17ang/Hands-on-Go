package main

import (
	"fmt"
)

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
	var alex person
	alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 5300,
		},
	}
	fmt.Println(alex)

	alexPointer := &alex
	alexPointer.updateName("Jimmy")
	alex.print()

	alexPointer.updateName("Neo")
	alex.print()

}

//Avoiding pass by value
//Using pass by reference

func (pt *person) updateName(newFirstName string) {
	(*pt).firstName = newFirstName
}

func (pt person) print() {
	fmt.Printf("%+v\n", pt)
}
