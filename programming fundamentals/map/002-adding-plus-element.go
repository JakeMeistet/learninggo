package main

import "fmt"

func main() {
	m := map[string]int{
		"Jake": 17,
		"Heidi": 19,
	}
		fmt.Println(m)

		fmt.Println(m["Jake"])

		fmt.Println(m["Test"])

	    if v, ok := m["Test"]; ok {
	    	fmt.Println(v)
		} else {
			fmt.Println("Doesn't exist")
		}

	    m["Todd"]= 33

	    for k, v := range m {
	    	fmt.Println(k, v)
		}

}

