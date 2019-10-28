package main

import (
	"fmt"
)

func main() {

	type person struct {
		firstName   string
		lastName    string
		favFlavours []string
	}

	p1 := person{
		firstName: "Jake",
		lastName:  "Bailey",
		favFlavours: []string{
			"Vanilla",
			"Chocolate",
		},
	}

	p2 := person{
		firstName: "James",
		lastName:  "Bond",
		favFlavours: []string{
			"Strawberry",
			"Cookie Dough",
		},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}

}
