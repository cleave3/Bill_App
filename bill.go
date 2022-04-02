package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//reciever functions
func (bill *Bill) format() string {
	fs := "Bill breakdowm: \n"

	var total float64 = 0 + bill.tip

	for k, v := range bill.items {
		fs += fmt.Sprintf("%-20v .... $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-20v .... $%0.2f\n", "Tip:", bill.tip)
	fs += fmt.Sprintf("%-20v .... $%0.2f\n", "Total:", total)

	return fs
}

// update Tip
func (bill *Bill) updateTip(tip float64) {
	bill.tip = tip
}

// add item to bill
func (bill *Bill) addItem(name string, price float64) {
	bill.items[name] = price
}

// save bill
func (bill *Bill) save() {
	data := []byte(bill.format())

	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
