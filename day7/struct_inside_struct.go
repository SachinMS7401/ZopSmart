package main

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Println("full address: %s ,%s", a.city, a.state)
	fmt.Println(a)
}
func (p person) fulldetails() {
	fmt.Println("full details:name:%s %s, address: %s, %s ", p.firstName, p.lastName, p.city, p.state)
	fmt.Println(p)
}

type person struct {
	firstName string
	lastName  string
	address
	age int
}

func main() {
	//gg := address{"fd", "dd"}

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
		age: 21,
	}
	fmt.Println(p.firstName)

	p.fullAddress() //accessing fullAddress method of address struct

	p.fulldetails()

}
