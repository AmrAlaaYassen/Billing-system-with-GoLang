package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Create new Bill
func createNewBill(name string) bill {
	bill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.00,
	}

	return bill
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for item, price := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", item+":", price)
		total += price
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "Tip: ", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total: ", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save () { 
	data := []byte (b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err!=nil { 
		panic(err)
	}


	fmt.Println(b.format())
}