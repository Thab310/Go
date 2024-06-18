package main

import "fmt"

func main () {
	//eg.1
	x := 0
	for x < 5 {
		fmt.Println("The value of x is:", x)
		x ++
	}

	//eg.2
	for i :=0; i < 5; i++ {
		fmt.Println("The value of i is:", i)
	}

	//eg.3
	names := []string {"Thabelo", "Kern", "Thulani", "Wendell"}

	for i :=0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	dataSet := []string {"car", "bike", "plane", "boat"}

	for x := 0; x < len(dataSet); x++{
		fmt.Printf("%v is number %v \n",dataSet[x],x)
	} 


	//eg.4
	for index, value := range names {
		fmt.Printf("The position at index %v is %v", index, value)
	}

	for _, value := range names {
		fmt.Println("The value is: \n",value)
	}

}