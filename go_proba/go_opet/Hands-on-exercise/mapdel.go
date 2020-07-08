package main

import "fmt"

func main() {

	m := map[string]int{
		"Gogi":  13,
		"Miki":  1301,
		"Kajo":  23,
		"Vjeko": 44,
	}
	fmt.Println(m)

	/*	delete(m, "Vjeko")
		fmt.Println(m)*/

	fmt.Println(m)
	mdel(m, "Vjeko")
	fmt.Println(m)

	if v, ok := m["Miki"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Miki")
	}
	fmt.Println(m)
}

func mdel(mp map[string]int, ent string) {
	for k, v := range mp {
		//fmt.Println(k, v)
		if k == ent {
			delete(mp, ent)
			//fmt.Println(ent, "  deleted!")
			fmt.Println(k, v, "deleted")
		} else {
			fmt.Println(ent, " not found inside map!")
			break
		}
	}
}
