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

}
