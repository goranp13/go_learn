package main

import "fmt"

func main() {

	sl := []int{12, 34, 23, 54, 6, 67, 26, 99, 65}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", sl)
}
