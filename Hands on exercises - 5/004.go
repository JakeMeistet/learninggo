package main

import "fmt"

func main() {

	annonymous := struct {
		fName     string
		friends   map[string]int
		favDrinks []string
	}{
		fName: "Jake",
		friends: map[string]int{
			"Heidi": 1,
			"Ben":   2,
		},
		favDrinks: []string{
			"Cream Soda",
			"Coca Cola",
		},
	}

	fmt.Println(annonymous)
	fmt.Println(annonymous.fName)
	fmt.Println(annonymous.friends["Heidi"])
	fmt.Println(annonymous.favDrinks[0], "\n")

	for _, v := range annonymous.favDrinks {
		fmt.Println(v)
	}
	fmt.Println()
	for k, v := range annonymous.friends {
		fmt.Println(k, v)
	}

}
