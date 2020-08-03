package main

import "fmt"

func main() {
	n1 := 1
	n2 := 3
	defer fmt.Println(n1 + n2)
	fmt.Println("n1 + n2 = ")

	defer foo()

}

func foo() {
	defer func() {
		fmt.Println("foo defer ran")
	}()
	fmt.Println("Foo ran")
}
