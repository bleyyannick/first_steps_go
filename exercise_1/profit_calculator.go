package main

import (
	"fmt"
)

func main() {
	var revenues float64
	var expenses float64
	var taxRate float64

	revenues = outputString("Revenues:")

	expenses = outputString("expenses:")

	taxRate = outputString("taxRate:")

	ebt, eat, ratio := calculateProfit(revenues, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(eat)
	fmt.Println(ratio)

}

func outputString(placeholder string) float64 {
	var result float64
	fmt.Print(placeholder)
	fmt.Scan(&result)
	return result
}

func calculateProfit(revenues, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenues - expenses
	eat = ebt - (revenues * taxRate / 100)
	ratio = ebt / eat
	return

}
