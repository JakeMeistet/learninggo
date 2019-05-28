package main

import "fmt"

var x int
var y float64

func main() {
	x = 42
	y = 42.34534

	fmt.Println(x,y)
	fmt.Printf("%T\t%T\t", x, y)
}
