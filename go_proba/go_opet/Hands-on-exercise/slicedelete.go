package main

import "fmt"

func main() {
	x := []int{13, 1, 1976, 29, 5}
	fmt.Println(x)
	x = append(x, 1976, 30, 10, 27, 6)
	fmt.Println(x)

	y := []int{12, 34, 57, 24, 2}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
