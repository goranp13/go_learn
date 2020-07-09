package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond_james`:       []string{`Shaken, not stirred`, `Martinis`, "Women"},
		`moneypenny_miss,`: []string{`James Bond`, `Literature`, `Computer science`},
		`no_dr`:            []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)

	m[`Tomo`] = []string{`Skladištar`, `Zajebava`, `Ne sluša`}

	delete(m, `no_dr`)

	for k, v := range m {
		fmt.Println("This is record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
