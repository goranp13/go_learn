package main

import (
	"log"
	"os"
	"time"
)

type Invoice struct {
	BrRn    int
	IznosRN int
}

var t = time.Now()
var d = t.Format("2006-01-02")

func main() {

	podaci := Invoice{
		BrRn:    2,
		IznosRN: 10,
	}

	invoicelog(podaci.BrRn, podaci.IznosRN)

}

func invoicelog(x int, y int) {
	file, e := os.OpenFile("invoice.txt "+d, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if e != nil {
		log.Fatalln("failed")
	}

	defer file.Close()

	log.SetOutput(file)

	//log.Println("Broj računa: " + x + "Iznos računa: " + y)
	log.Printf("\nRačun broj:%+v \t Iznos računa:%+v kn\n\n", x, y)
}
