package main

import "fmt"

func main() {
	age := 21
	drink := true
	driving := age == 21 && drink == false

	switch driving {
	case (age == 21 && drink != true):
		fmt.Println("Go baby, drive :-)!")
	case age != 21 && drink != true:
		fmt.Println("Sorry you're underaged!")
	}
}
