package main

import (
	"fmt"
)

func main(){
	age := 36

	fmt.Println(age <= 50)
	fmt.Println(age >= 40)
	fmt.Println(age == 35)
	fmt.Println(age != 47)

	if age < 35 {
		fmt.Println("Age is less than 35")
	} else if age > 40{
		fmt.Println("You're above 40 years")
	} else {
		fmt.Println("You're between 35 and 40 years")
	}

	dataSet := []string {"car", "bike", "plane", "boat"}

	for index,value := range dataSet {
		if index == 1 {
			fmt.Printf("continuing at pos %v\n",index)
			continue //go back to the loop forget about the rest of the code
		}

		fmt.Printf("The value at pos %v is %v \n",index, value)
		
	}

	for i,v := range dataSet{
		if i ==1 {
			fmt.Println("continuing at pos ", i)
			continue
		}
		if i > 2 {
			fmt.Printf("the value at pos %v is %v \n", i,v)
			break //break out of the loop completely
		}
	}

}
