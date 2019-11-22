package main

import "fmt"

type person struct {
	first string
	last string
}

type agent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code... }

func (s agent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	p1 := agent{
		person: person{
			"Jake",
			"Bailey",
		},
		ltk: true,
	}

	fmt.Println(p1)
	p1.speak()

}
