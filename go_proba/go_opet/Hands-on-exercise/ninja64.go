package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	o1 := person{
		first: "Goran",
		last:  "Pozder",
		age:   44,
	}

	o1.speak()

}

func (s person) speak() {
	fmt.Println("I am:", s.first, s.last, "and I am ", s.age, "years old.")
}
