package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++  {
		// %+q ASCII - godoc.org/fmt
		fmt.Printf("Denary: %d\t\t ASCII: %+q\n", i, i)
	}
}
