package main

import (
	"fmt"
)

func main() {
	x := 2
	y := 4

	if x == y {
		fmt.Println("x is equal to y")
	} else if x < y {
		fmt.Println("x is less than y")
	}else if y < x {
		fmt.Println("y is less than x")
	}else if y > x {
		fmt.Println("y is greater than x")
	}else {
		fmt.Println("x is greater than y")
	}
}
