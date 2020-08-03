package main

import "fmt"

func main() {
	L1 := struct {
		name    string
		obitelj map[string]int
		razred  []string
	}{
		name: "Lena",
		obitelj: map[string]int{
			"Tata": 100,
			"Mama": 100,
			"Vid":  100,
		},
		razred: []string{
			"1.razred",
			"2.razred",
			"3.razred",
		},
	}

	fmt.Println(L1)
	fmt.Println(L1.name)
	fmt.Println(L1.name)
	fmt.Println(L1.razred)

	for i, v := range L1.obitelj {
		fmt.Println("Obitelj: ", i, v)
	}

	for _, val := range L1.razred {
		fmt.Println(val)
	}
}
