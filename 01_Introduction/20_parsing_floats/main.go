package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add items, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}

		b.addItems(name, p)

		fmt.Println("items added - ", name, price)
		promtOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promtOptions(b)
		}

		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promtOptions(b)
	case "s":
		fmt.Println("You chose to save the bill", b)
	default:
		fmt.Println("That was not a valid option...")
		promtOptions(b)
	}
}

func main() {
	myBill := createBill()

	promtOptions(myBill)
}
