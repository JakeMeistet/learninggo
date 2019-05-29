package main

import "fmt"

func main() {
	x := 2
	y := 4
	switch {
	case x == y:
		fmt.Println("x is equal to y")
	case y > x:
		fmt.Println("y is greater than x")
	default:
		fmt.Println("Nothing is true")
	}
}