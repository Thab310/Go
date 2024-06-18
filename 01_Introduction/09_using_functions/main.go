package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n",n)
}

func sayBye(n string){
	fmt.Printf("Good bye %v \n", n)
}


func cycleNames(n []string,f func(string) ){
	for i,v := range n {
		fmt.Printf("The index %v is at position %v \n", i,v)
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main(){
	sayGreeting("Thabelo")
	sayGreeting("Bono")
	sayBye("Lili")
	sayBye("Mike")

	mySlice := []string{"car","bike","plane"}
	cycleNames(mySlice, sayGreeting)

	a1 := circleArea(10.4)
	a2 := circleArea(5)

	fmt.Println(a1,a2)
	fmt.Printf("cirle a1 is %0.3f and circle 2 is %0.3f \n",a1,a2)
}