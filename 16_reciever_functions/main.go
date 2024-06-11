package main

import "fmt"

func main() {
	myShop := newShop("Thabelo's shop")

	myShop.format()
	fmt.Println(myShop.format())
}
