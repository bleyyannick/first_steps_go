package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64 // := recommended way of assigining an infered type to a variable
	var expectedReturnRate float64
	var years float64

	outputText("Expect return rate:") // used for writing messages for users
	fmt.Scan(&expectedReturnRate)     // & sign is used for storing inputs user

	outputText("Years:")
	fmt.Scan(&years)

	outputText("Invest Amount:")
	fmt.Scan(&investmentAmount) // pointers used to assign value entered by user

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	formattedBV := fmt.Sprintf("Future value: %1.f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation)\n: %1.f", futureRealValue)
	fmt.Println(formattedBV)
	fmt.Println(formattedRFV)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, realFutureValue float64) {
	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue
}

func outputText(text string) {
	fmt.Print(text)
}
