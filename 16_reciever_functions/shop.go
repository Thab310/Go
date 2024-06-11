package main

import "fmt"

type shop struct {
	name  string
	items map[string]float64
	price float64
}

func newShop(n string) shop {
	s := shop{
		name:  n,
		items: map[string]float64{"catfish": 18.99, "hake": 35.99, "salmon": 65.99},
		price: 0,
	}

	return s
}

// format function can only be called by the shop objects
func (s shop) format() string {
	fs := "shop breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range s.items {
		fs += fmt.Sprintf("%-30v ...R%v \n", k+":", v)
		total += v
	}

	//total
	fs += fmt.Sprintf("%-30v ...R%0.2f", "total:", total)

	return fs
}
