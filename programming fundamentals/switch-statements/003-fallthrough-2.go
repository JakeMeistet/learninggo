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
		fallthrough
	case 7 == 9:
		fmt.Println("No")
		fallthrough
	case 11 == 14:
		fmt.Println("No")
		fallthrough
	case 15 == 15:
		fmt.Println("Yes")
	}
}