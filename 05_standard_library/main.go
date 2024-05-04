package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	warning := "Hi please beware of the dog!"

	fmt.Println(strings.Contains(warning, "Hi")) //returns true or false 
	fmt.Println(strings.ReplaceAll(warning, "Hi", "Bonjour"))
	fmt.Println(strings.ToUpper(warning))
	fmt.Println(strings.Index(warning, "p"))
	fmt.Println(strings.Split(warning, " ")) //split into slice

	marks := []int {24,45,12,97,51,67,1}

	sort.Ints(marks)
	fmt.Println(marks)

	names := []string {"Thabelo", "Mike", "Daniel", "Steve"}

	index := sort.SearchInts(marks, 12)
	
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(index)
	fmt.Println(sort.SearchStrings(names, "Steve"))
}	