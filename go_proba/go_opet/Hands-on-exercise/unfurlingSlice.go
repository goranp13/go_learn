package main

import "fmt"

func main() {
	xi := []int{2, 4, 5, 67, 3, 2, 9, 8, 6}
	x := sum(xi...)
	fmt.Println("The total is: ", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index", i, "we add", v, "to the total", "which is: ", sum)
	}
	fmt.Println("The  total is: ", sum)
	return sum
}
