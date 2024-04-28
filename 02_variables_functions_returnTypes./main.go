package main

import "fmt"
//Note!! we can initialize variables outside of the function but we cannot assign values to them
//Eg. var pilot string = "Steve" this is not valid because 

var mechanic string
var driver string

func main() {
	mechanic = "Thabelo"
	driver = "John"

	fmt.Println(mechanic)
	fmt.Println(driver)


	var car string = "bmw"
	fmt.Println(car) 

	var boat = "kawasaki"
	fmt.Println(boat)

	plane := "gulfstream"
	fmt.Println(plane)

	fmt.Println(color())

	car = newCar()
	fmt.Println(newCar())
}

func color() string {
	return "blue"
}

func newCar() string {
	return "Toyota"
}