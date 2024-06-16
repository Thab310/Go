package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(n string) bill {
	b := bill{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

// format function can only be called by the bill objects
func (b bill) format() string {
	fs := fmt.Sprintf("%v \n", b.name)
	fs += "bill: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-30v ...R%v \n", k+":", v)
		total += v
	}

	//ADD TIP
	fs += fmt.Sprintf("%-30v ...R%v \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-30v ...R%0.2f", "total for 3:", total+b.tip)

	return fs
}
