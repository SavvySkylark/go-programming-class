package main

import "fmt"

func main() {
	m := map[string][]string{
		"james_bond":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(m, "no_dr")

	for _, p := range m {
		for i, phrase := range p {
			fmt.Printf("%v, %v\n", i, phrase)
		}
	}
}
