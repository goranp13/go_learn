package main

import "fmt"

func main() {

	g := func(xi []int) int {
		for _, v := range xi {
			fmt.Println(v)
		}
		return xi[0]
	}
	x := cb(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)

}

func cb(cbl func(xi []int) int, ii []int) int {
	n := cbl(ii)
	n++
	return n
}
