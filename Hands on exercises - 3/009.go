package main

import "fmt"

func main() {
	favSport := "Football"
	switch favSport {
	case "Football":
		fmt.Println("will print")
	case "Hockey":
		fmt.Println("won't print")
	default:
		fmt.Println("Nothing is true")
	}
}
