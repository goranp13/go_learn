package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("_____________________")

	fmt.Printf("Name:%v \tLastname:%v \tAge:%v \n", p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(sa1.person)
}
