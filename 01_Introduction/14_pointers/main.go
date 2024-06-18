package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}


func main(){
	name := "thabelo"


	fmt.Println("The memory address of name is: ", &name) //pointer to the memory location

	m:= &name //pointer
	fmt.Println("memorty address:", m)
	fmt.Println("Value at memory address:", *m) //value of pointer


	updateName(m)
	fmt.Println(name)

	fmt.Println(m)

}

/*
|---name--- | ---m---|
|   0x001   |  0x002 |
|-----------|--------|
|"Thabelo"  | p0x001 |
|-----------|--------|
*/