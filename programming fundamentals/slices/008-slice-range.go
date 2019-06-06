package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9, 42}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

}
