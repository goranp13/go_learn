package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {

	Lena := person{
		firstName: "Lena",
		lastName:  "Pozder",
		favoriteIceCreamFlavors: []string{
			"Vanila",
			"Strawberry",
			"Chocolate",
		},
	}

	Vid := person{
		firstName: "Vid",
		lastName:  "Pozder",
		favoriteIceCreamFlavors: []string{
			"Apple",
			"Strawberry",
			"Limone",
		},
	}

	fmt.Println(Lena)
	fmt.Println(Vid)

	for i, v := range Lena.favoriteIceCreamFlavors {
		fmt.Println("Lena's favorite ic cream flavors are: ", i+1, v)
	}

	fmt.Println("")

	for i, v := range Vid.favoriteIceCreamFlavors {
		fmt.Println("Vid's favorite ic cream flavors are: ", i+1, v)
	}

}
