package main

import "fmt"

func main() {

	sum(2, 3, 45, 2, 4, 65, 43, 2, 3, 67)

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
