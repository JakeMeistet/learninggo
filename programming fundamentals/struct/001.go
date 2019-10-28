package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2, "\n")

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
