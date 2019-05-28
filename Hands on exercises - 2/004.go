package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("Binary: %b\t\t Denary: %d\t\t Hex: %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("Binary: %b\t\t Denary: %d\t\t Hex: %#x\n", y, y, y)
}