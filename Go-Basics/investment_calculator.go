package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	//var investmentAmount = 1000
	// var investmentAmount float64 = 1000
	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount, years float64
	// years := 10.0
	// var expectedReturnRate = 5.5
	expectedReturnRate := 5.5
	// var years = 10
	// var years float64 = 10
	// investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	
	// fmt.Println(futureRealValue)

	// fmt.Printf("The future value of the investment is: %v\n", futureValue)
	// fmt.Printf("The future real value of the investment (adjusted for %v%% inflation) is: %.v\n", inflationRate, futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value (adjusted for %.2f%% inflation): %.2f\n", inflationRate, futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFRV)
	// fmt.Print(formattedFV, formattedFRV)

}

// func outputText(text, text2 string)  
func outputText(text string)  {
	fmt.Print(text)	
}

func calculateFutureValue(principal, rate, time float64) (float64, float64) {
	fv := principal * math.Pow(1 + rate/100, time)
	rfv := fv / math.Pow(1 + inflationRate/100, time)
	return fv, rfv
}

// func calculateFutureValue(principal, rate, time float64) (fv float64, rfv float64) {
	// fv = principal * math.Pow(1 + rate/100, time)
	// rfv = fv / math.Pow(1 + inflationRate/100, time)
	// return
	// return fv, rfv
// }