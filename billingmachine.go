package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var orders string
	var quantity int
	var sales = []int{}
	var counter int
	counter = 0
	for true {
		fmt.Println("MENU")
		fmt.Println("-----Availble Dishes for you-----")
		fmt.Print("C : coffee : 40rs\n D : dosa : 80rs\n T : tomato soup : 20rs\n I : idli : 20rs\n V : vada : 25rs\n B : bhature &chhole : 30rs\n P : paneer pakoda : 30rs\n M: manchurian : 90rs\n H : hakka noodle : 70rs.\n F : French fries : 30rs\n J : jalebi : 30rs\n L : lemonade : 15rs\n S : spring roll : 20rs\n")
		fmt.Println("Press END For closing the day.")
		fmt.Println("place the order: ")

		fmt.Scan(&orders)
		orders = strings.ToUpper(orders)

		if orders == "END" {
			total_income(sales)
		} else {
			fmt.Scan(&quantity)
		}
		var Total_bill int = Total_bill(quantity, orders)
		fmt.Println("Please wait for the bill to be printed...")
		fmt.Println("Total bill: ", Total_bill)
		fmt.Println("Your billing is successfully")
		sales = append(sales, Total_bill)
		counter++

	}
}

func total_income(sale []int) {
	var total int
	total = 0

	for i := 0; i < len(sale); i++ {
		total = sale[i] + total
	}
	fmt.Println("Total Income is : ", total)
	os.Exit(0)
}

func Total_bill(quantity int, order string) int {
	m := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	var your_bill int
	your_bill = quantity * m[order]
	return your_bill
}
