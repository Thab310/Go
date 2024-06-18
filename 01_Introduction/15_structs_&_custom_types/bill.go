// This is a custom structure for any bill object
package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"catfish": 18.99, "hake": 35.99, "salmon": 65.99},
		tip:   0,
	}

	return b
}
