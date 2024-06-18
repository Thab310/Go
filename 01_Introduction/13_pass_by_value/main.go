package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["pie"] = 12.98
}
func main(){
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "thabelo"

	name = updateName(name)
	fmt.Println(name)

	//group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":2.3,
		"cake": 45,
	}

	updateMenu(menu)
	fmt.Println(menu)
}