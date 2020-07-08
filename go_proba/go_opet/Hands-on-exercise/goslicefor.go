package main

import "fmt"

func main() {

	x := []int{13, 1, 1976, 29, 5}
	fmt.Println(len(x))
	fmt.Println("1st index number:\t", x[0])
	fmt.Println("2nd index number:\t", x[1])
	fmt.Println("3rd index number:\t", x[2])
	fmt.Println("4th index number:\t", x[3])
	fmt.Println("5th index number:\t", x[4])

	for i, v := range x {
		fmt.Printf("%v Index:\t value:%v\n", i, v)
	}

}
