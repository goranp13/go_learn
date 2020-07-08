package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("Go skii!")
	case "swimming":
		fmt.Println("Go swim!")
	case "surfing":
		fmt.Println("Yeah, surfing!")
	case "windsurfing":
		fmt.Println("You are fucked!")
	}

}
