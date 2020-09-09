package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 9, 1, 7, 45, 8, 4, 3}
	xs := []string{"Gogi", "Lena", "Vid", "Nata", "mama", "kćer", "tata", "sin"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("---------------------------------------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
