package main

import "fmt"

type Jake int

var x Jake

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
}
