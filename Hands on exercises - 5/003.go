package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Coffee",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(t.doors)

	fmt.Println(s)
	fmt.Println(s.doors)

}
