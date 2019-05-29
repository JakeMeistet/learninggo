package  main

import "fmt"

func main() {
	switch "Hello" {
	case "Test", "Hello", "TesT":
		fmt.Println("Hello World")
	default :
		 fmt.Println("Default")
	}

}