package main

import "fmt"

type shop struct {
	name  string
	items map[string]float64
	tip   float64
}

func newShop(n string) shop {
	s := shop{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}

	return s
}

//update tip
func (s *shop) updateTip(tip float64) {
	s.tip = tip
}

// add item to shop
func (s *shop) addItems(name string, price float64) {
	s.items[name] = price
}

// format function can only be called by the shop objects
func (s shop) format() string {
	fs := fmt.Sprintf("%v \n", s.name)
	fs += "menu: \n"
	var total float64 = 0

	//list items
	for k, v := range s.items {
		fs += fmt.Sprintf("%-30v ...R%v \n", k+":", v)
		total += v
	}

	//ADD TIP
	fs += fmt.Sprintf("%-30v ...R%v \n", "tip:", s.tip)

	//total
	fs += fmt.Sprintf("%-30v ...R%0.2f", "total for 3:", total+s.tip)

	return fs
}
