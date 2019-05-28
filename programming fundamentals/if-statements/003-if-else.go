package main

import "fmt"

func main() {
	x := 41
	if x == 42 {
		fmt.Println("Is 42")
	} else if x == 41 {
		fmt.Println("Is 41")
	} else {
		fmt.Println("Neither")
	}
}
