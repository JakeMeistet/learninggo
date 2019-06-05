package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken", "Not", "Stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hello,", "James"}

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Printf("\nRecord: %v", i)
		for j, val := range xs {
			fmt.Printf("\n Index position: %v \t Value: \t %v", j, val)
		}
	}

}