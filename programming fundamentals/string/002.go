package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	// slice of byte
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Println(i,v)
	}
}
