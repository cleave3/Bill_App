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

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to superstore")

	name, _ := getInput("What's your name ? ", reader)
	name = strings.TrimSpace(name)

	if name == "" {
		createBill()
	}

	bill := newBill(name)

	fmt.Printf("A new tab created - %v's bill \n", bill.name)

	return bill
}

func promptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("What do you what to do ?\na - add item\ns - save bill\nt - add tip\n", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(bill)
		}

		bill.addItem(name, p)

		fmt.Println("Item added ", name, price)

		promptOptions(bill)

	case "s":
		bill.save()
		fmt.Println("you saved the bill to file: ", bill.name)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(bill)
		}
		bill.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOptions(bill)
	default:
		fmt.Println("Invalid option")
		promptOptions(bill)
	}

}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
