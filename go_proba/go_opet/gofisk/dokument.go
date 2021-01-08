package main

import (
	"time"
)

//Dokument struktura
type Dokument struct {
	Naziv string
	Body  []byte
}

var t = time.Now()
var d = t.Format("2006-01-02 15-04-05")

func main() {
	d1 := &Dokument{Naziv: "Račun", Body: []byte("iznos računa 12 kn, broj računa 12231234")}
	//d1 := &Dokument{Naziv: "Report", Body: []byte("iznos računa za 1 mjesec iznosi 1773 kn kn, broj reporta 1")}
	d1.save()
}


func (p *Dokument) save() error {
	filename := p.Naziv + "_" + d + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

