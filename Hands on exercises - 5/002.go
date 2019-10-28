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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)
		}
	}
}
