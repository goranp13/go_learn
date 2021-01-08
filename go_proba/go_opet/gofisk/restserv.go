package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Dokument struct {
	Naziv string `json:"Naziv"`
	Body  string `json:"Body"`
}

//Dokumenti računi i reporti
type Dokumenti []Dokument

func main() {

	/* 	Dokumenti = []Dokument{
		Dokument{Naziv: "Račun br.1", Body: "Račun za parkiranje vozila, iznos 12 kn"},
		Dokument{Naziv: "Report Ukupno utržak", Body: "Izvještaj o ukupnom prometu"},
	} */

	//	Dokumenti = []testDokumenti{}

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	//myRouter.HandleFunc("/dokumenti", sviDokumenti).Methods("GET")
	//myRouter.HandleFunc("/dokumenti", sviDokumenti)
	myRouter.HandleFunc("/dokumenti", testDokumenti).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func sviDokumenti(w http.ResponseWriter, r *http.Request) {
	dokumenti := Dokumenti{
		Dokument{Naziv: "RN1", Body: "Ovo je testni račun bez iznosa"},
	}
	fmt.Println("Endpoint: svi dokumenti")
	json.NewEncoder(w).Encode(dokumenti)

}

func testDokumenti(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test POST testni Dokumenti radi")
	json.NewEncoder(w).Encode(dokumenti)
}

/* func returnSviDokumenti(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Prikaz svih dokumenata: popis dokumenata")
	json.NewEncoder(w).Encode(Dokumenti)
}
*/
