package main

import "fmt"

func main() {
	foo()
}

// func (r receiver)  identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Function text")
	bar("Jake")
	s1 := woo("Heidi")
	fmt.Println(s1)
}

// everything in Go is pass ByVal
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
} 