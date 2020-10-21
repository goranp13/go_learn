package main

import "fmt"

func main() {

	poruke := make(chan string, 3)
	kan := make(chan int)

	poruke <- "ovo je"
	poruke <- "proba"

	fmt.Println(<-poruke)
	fmt.Println(<-poruke)

	poruke <- "rada buffer kanala"
	fmt.Println(<-poruke)

	x := 2
	y := 3

	go zbr(x, y, kan)
	fmt.Println(<-kan)
}

func zbr(x int, y int, ch chan<- int) {
	s := x + y
	ch <- s
}
