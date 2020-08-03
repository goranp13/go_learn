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

	car1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: false,
	}

	car2 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.color)
	fmt.Println(car2.luxury)

}
