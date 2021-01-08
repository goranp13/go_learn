package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var t = time.Now()
var d = t.Format("2006-01-02 15-04-05")

//Invoice struktura računa
type Invoice struct {
	BrRn    int
	IznosRN int
}

func main() {

	p := t.Format("15:04:05")

	fmt.Println(p)

	podaci := Invoice{
		BrRn:    1,
		IznosRN: 30,
	}

	file, _ := json.MarshalIndent(podaci, " ", " ")

	_ = ioutil.WriteFile("invoice.json "+d, file, 0644)

}
