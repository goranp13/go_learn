package main

import "fmt"

func main() {

	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)

	//ili
	fmt.Println(bar()())

	//	i := x()
	//	fmt.Println(i)
	fmt.Println(x())
}

func foo() string {
	s := "Hello from foo"
	return s
	//ili samo return "Hello from foo "
}

func bar() func() int {
	return func() int {
		return 451
	}
}
