package main

import "fmt"

func main()  {
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"cocolate pudding": 14.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu{
		fmt.Println(k,"-",v)
	}

	// ints as key type
	phonebook := map[int]string{
		06100000: "Mike",
		07100000: "Buhle",
		07200000: "Thabelo",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[06100000])

	phonebook[06100000]="Laden"
	fmt.Println(phonebook)

}