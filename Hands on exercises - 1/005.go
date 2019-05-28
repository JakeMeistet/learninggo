package main

import "fmt"

// UNDERLYING TYPE of int is int
// Therefore the UNDERLYING TYPE of Jake is int

type Jake int

var x Jake
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
