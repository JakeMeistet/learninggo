package main

import "fmt"

func main() {
	x := 2002
	for {
		if x > 2019 {
			break
		}
		fmt.Println(x)
		x++
	}
}
