package main

import "fmt"

func main() {
	plain := []int{1, 0, 1, 1, 0, 1, 0, 0}
	fmt.Println("Plain:\t\t", plain)
	key := []int{0, 1, 1, 1, 0, 1, 1, 0}
	fmt.Println("Key:\t\t", key)
	encrypted := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 8; i++ {
		if plain[i] == key[i] {
			encrypted[i] = 0
		} else {
			encrypted[i] = 1
		}
	}
	fmt.Println("Encrypted:\t", encrypted)
}
