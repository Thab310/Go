package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string //initialize the variable as a slice but not assigned a value
	
	for _, v := range names {
		fmt.Println(v[:1])
		initials = append(initials, v[:1]) //On every iteration take value of index 0 within the initials slice so we have a slice within a slice 
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func main()  {
	fn,sn := getInitials("thabelo ramabulana")

	fmt.Println(fn,sn)
	
}