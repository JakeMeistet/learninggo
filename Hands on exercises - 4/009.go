package main

import "fmt"

func main() {
	m := map[string][]string {
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["bailey_jake"] = []string{"Coffee", "Food", "Computer Science"}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}
}
