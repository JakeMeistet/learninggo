package main

import "fmt"

func main() {
	jb := []string{"Jake", "Bailey", "Chocolate", "Cookie"}
	fmt.Println(jb)

	test := []string{"test", "test2", "test3", "test4"}
	fmt.Println(test)

	xp := [][]string{jb, test}
	fmt.Println(xp)
}
