package  main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("No")
	case 2 == 4:
		fmt.Println("No")
	case 3 == 3:
		fmt.Println("Yes")
		fallthrough
	case 4 == 4:
		fmt.Println("Yes")
	}
}