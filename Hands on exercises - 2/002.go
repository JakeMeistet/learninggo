package main

import "fmt"

func main() {
	a := 67
	b := 98

	g := a == b
	h := a <= b
	i := a >= b
	j := a != b
	k := a < b
	l := a > b

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

}
