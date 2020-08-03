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

func (s djelatnik) reci() {
	fmt.Println("Ja sam: ", s.ime, s.prezime)
}

func main() {

	d1 := djelatnik{
		osoba: osoba{
			ime:     "Goran",
			prezime: "Pozder",
		},
		radDozvola: true,
	}

	d1.reci()

}
