package main

import "fmt"

func main() {
	f := gogi()

	fmt.Println(f())

}

func gogi() func() int {
	return func() int {
		return 44
	}
}

