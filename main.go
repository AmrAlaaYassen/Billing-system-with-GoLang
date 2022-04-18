package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promot string, r *bufio.Reader) (string, error) {
	fmt.Print(promot)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t - add tip):  ", reader)

	switch opt {
	case "a":
		fmt.Println("you choose to add an item")
		name, _ := getInput("Please enter item name: ", reader)
		price, _ := getInput("Please enter item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil { 
			fmt.Println("The price must be a number")
		}else { 
			b.addItem(name, p)
			fmt.Println("Item add - ", name, " : ", price)
		}
		promptOptions(b)
	case "s":
		fmt.Println("You choose to save the bill") 
		b.save()	
	case "t":
		fmt.Println("You choose to add a tip")
		tip, _ := getInput("Please enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil { 
			fmt.Println("Tip must be a number")
		} else { 
			b.updateTip(t)
		  println("Tip added - ", tip)
		}

		promptOptions(b)
	default:
		fmt.Println("Incorrect option ...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create new bill name: ", reader)

	new_bill := createNewBill(name)
	fmt.Println("Created Bill - ", new_bill.name)

	return new_bill
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
