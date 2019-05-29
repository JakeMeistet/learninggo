package  main

import "fmt"

func main() {
	switch "Hello" {
	case "Test":
		fmt.Println("No")
	case "Hello":
		fmt.Println("Hello World")
	default :
		 fmt.Println("Default")
	}

}