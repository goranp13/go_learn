package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("My first func expression")
	}

	f()

	g := func(x int) {
		fmt.Println("Big brother ", x)
	}
	
	g(44)
}
