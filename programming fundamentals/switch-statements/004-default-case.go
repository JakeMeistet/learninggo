package  main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("No")
	case 2 == 4:
		fmt.Println("No")
	default :
		 fmt.Println("Default")
	}

}