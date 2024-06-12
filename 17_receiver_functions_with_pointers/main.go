package main

import "fmt"

func main() {
	myShop := newShop("Welcome to Thabelo's butchery")

	myShop.updateTip(10)
	myShop.addItems("steak", 76)
	myShop.addItems("chicken", 56)
	myShop.addItems("salmon", 80.99)
	fmt.Println(myShop.format())
}
