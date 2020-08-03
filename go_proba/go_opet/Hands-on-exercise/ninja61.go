package main

import "fmt"

func main() {

	G := foo()
	fmt.Println(G)

	P, p := bar()
	fmt.Println(P, p)

}

func foo() int {
	return 13
}

func bar() (int, string) {
	return 1, "Goran"
}
