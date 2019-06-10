package main

import "fmt"

type person struct {
	first string
	last  string
}

type secret_agent struct {
	person
	ltk bool
}

func main() {
	p1 := secret_agent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := secret_agent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("")
	fmt.Println(p1.first, p1.last, p1.ltk)
	fmt.Println(p2.first, p2.last, p2.ltk)

}
