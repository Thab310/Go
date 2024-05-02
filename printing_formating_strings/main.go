package main

import "fmt"


var age string

func main()  {
	name:= "Thabelo"
	age = "一百 years old"

	//Print
	fmt.Print("Hi bro")
	fmt.Print("Hi sis \n")
	fmt.Print("Hi everyone")

	//Println
	fmt.Println("Hi There")
	fmt.Println("Hi yall")
	fmt.Println("Hi my name is",name, "and I am",age)

	//Printf (Formatted strings) %_ format specifier
	fmt.Printf("Hi my name is %v and I am %v \n",name,age)
	fmt.Printf("Hi my name is %q and  I am %q \n", name,age) //places quotes around varaible....  does not work for numbers only works for strings
	fmt.Printf("Hi my name is of type %T and my age is of type %T \n",name,age)
	fmt.Printf("I am %f m tall \n", 1.71)
	fmt.Printf("I am %0.1f m tall \n", 1.71)
	fmt.Printf("I am %0.2f m tall \n", 1.71)

	//Sprintf (like "Printf", but saves the the formatted string in variable)
	var str = fmt.Sprintf("Hi my name is %v and I am %v \n", name, age)
	fmt.Println("My bio is:",str)
}

