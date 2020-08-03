package main

import "fmt"

type car struct {
	model string
	year  int
	state string
}

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	car1 := struct {
		model string
		year  int
		state string
	}{
		model: "BMW",
		year:  1987,
		state: "new",
	}

	car2 := car{
		model: "Mercedes",
		year:  2010,
		state: "used",
	}

	fmt.Println(p1)
	fmt.Println(car1)
	fmt.Println(car2)

}
