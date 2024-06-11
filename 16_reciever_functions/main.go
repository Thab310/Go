package main

import "fmt"

func main() {
	myShop := newShop("Welcome to Thabelo's butchery")

	fmt.Println(myShop.format())
}
