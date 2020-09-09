package main

import "fmt"

func main() {

	func(a int, b int) {
		fmt.Println(a * b)
	}(12, 45)

}
