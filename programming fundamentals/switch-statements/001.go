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
	case 4 == 4:
		// True but it wont print because the case before is also true
		fmt.Println("Yes")
	}
}