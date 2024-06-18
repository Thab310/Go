package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(n string) string {
	return englishHelloPrefix + n
}

func main() {
	fmt.Println(Hello("Thabelo"))
}
