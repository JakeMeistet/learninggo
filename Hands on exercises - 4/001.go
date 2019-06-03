package main

import "fmt"

func main() {
	x := [5]int{2,4,6,8,10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("Type: %T\n", x)
}
