package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func main() {

	djelatnik := person{
		name:    "Goran",
		surname: "Pozder",
		age:     44,
	}

	ime := "Gogi"
	//ime, _ := fmt.Scan()

	fmt.Println(djelatnik)
	changeMe(&djelatnik, ime)
	fmt.Println(djelatnik)

	var broj int
	fmt.Println("Unesi broj: ")
	fmt.Scanln(&broj)

	kanal := make(chan int)
	go mbytwo(broj, kanal)
	//fmt.Println(<-kanal)
	fmt.Printf("rezultat: %v\n", <-kanal)

	poruka := make(chan string)
	go func() { poruka <- "kanal poruka radi" }()

	por := <-poruka
	fmt.Println(por)

}

func changeMe(p *person, s string) {
	p.name = s
	//(*p).name = "Zdravko"
}

func mbytwo(p int, c chan<- int) {
	result := p * p
	c <- result
}
