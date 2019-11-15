package main

import "fmt"

func main() {
	foo()
}

// func (r receiver)  identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Function text")
	bar("Jake")
}

// everything in Go is pass ByVal
func bar(s string) {
	fmt.Println("Hello", s)
}
