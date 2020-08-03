package main

import "fmt"

type osoba struct {
	ime     string
	prezime string
}

type djelatnik struct {
	osoba
	radDozvola bool
}

type human interface {
	reci()
}

func (s djelatnik) reci() {
	fmt.Println("Ja sam: ", s.ime, s.prezime)
}

func (p osoba) reci() {
	fmt.Println("Ja sam: ", p.ime, p.prezime)
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {

	d1 := djelatnik{
		osoba: osoba{
			ime:     "Goran",
			prezime: "Pozder",
		},
		radDozvola: true,
	}

	p1 := osoba{
		ime:     "Miki",
		prezime: "Super",
	}

	d1.reci()
	p1.reci()
	bar(d1)
	bar(p1)
	//conversion

}
