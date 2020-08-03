package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("Anonxmus function ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(44)

}

func foo() {
	fmt.Println("Foo ran")
}
