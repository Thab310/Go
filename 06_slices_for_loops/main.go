package main

import (
	"fmt"
)

func main() {
	car := []string {"bmw","mercedez","toyota"} 
	var boat = []string {"kawasaki","honda", "lamborghini"}
	fmt.Println(car)
	fmt.Println(boat)

	vehicles := []string {car[len(car)-1], boat[len(boat)-1], plane()}
	fmt.Println(vehicles)

	for i, vehicle := range vehicles {
		fmt.Println(i, vehicle)
	}
}

func plane() string {
	return "gulfstream"
}