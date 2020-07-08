package main

import "fmt"

func main() {
	x := []int{13, 1, 1976, 29, 5}
	fmt.Println("sčicing a slice, print only what is up to index 1")
	fmt.Println(x[:1])
	fmt.Println("print what is between index 1 to 3")
	fmt.Println(x[1:3])

	for i, v := range x[2:4] {
		fmt.Println(i, v)
		fmt.Printf("%v index:\t value:%v\n", i, v)
	}
}
