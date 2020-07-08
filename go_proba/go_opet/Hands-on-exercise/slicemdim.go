package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocoalte", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hatzelnut"}
	fmt.Println(mp)
	//multi dimensional slice- (slice composed of sliuces)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
