package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["!Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("Barnabas is in the map", v)
	} else {
		fmt.Println("Barnabas is not in the map", v)
	}

	m["Todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 6, 5, 12, 4, 66}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
