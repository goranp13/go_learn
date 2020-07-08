package main

import "fmt"

func main() {
	var ar [5]int
	ar[0] = 1
	ar[1] = 2
	ar[2] = 3
	ar[3] = 4
	ar[4] = 5

	fmt.Println(ar)
	fmt.Printf("%T\n", ar)
	fmt.Println("-----------------------------------------")

	//teachers solution

	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
