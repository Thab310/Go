package main

import "fmt"

func main()  {
	//arrays
	var names = [4] string {"Thabelo", "Mike", "Tshikudo", "Makhadzi"}
	position := [4] int {1,2,3,4}
	names[2] = "laden"

	fmt.Println(names, len(names))
	fmt.Println(position, len(position))

	//slices
	forexPairs := []string {"EURZAR","EURAUD", "EURUSD", "GBPUSD" }
	fmt.Println(forexPairs)
	forexPairs[1]= "EURGBP"

	
	forexPairs = append(forexPairs, "USDZAR")
	fmt.Println(forexPairs)

	//slices ranges
	forexRangeOne := forexPairs[2:3]
	forexRangeTwo := forexPairs[2:]
	forexRangeThree := forexPairs[:2]
	fmt.Println(forexRangeOne)
	fmt.Println(forexRangeTwo)
	fmt.Println(forexRangeThree)


}