package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// => Show error message & exit if invalid input is provided
// - No negative numbers
// - Not 0
// 2) Store calculated results into file

var firstWrite = true
var resultFileName = "results.txt"

func main() {
	// var revenue, expenses, taxRate float64

	revenue, err1 := getUserInput("Enter the total revenue: ")
	expenses, err2 := getUserInput("Enter the total expenses: ")
	taxRate, err3 := getUserInput("Enter the tax rate (in %): ")
	if err1!= nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}
	ebt, profit, ratio := calculation(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)

	fmt.Printf("The profit ratio (EBT to Profit) is: %f\n", ratio)
	// writeResultsToFile(ratio)
	fmt.Printf("The profit before tax is: %f\n", ebt)
	// writeResultsToFile(ebt)
	fmt.Printf("The profit after tax is: %f\n", profit)
	// writeResultsToFile(profit)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid input. Please enter a positive number greater than zero.")
	}
	return userInput, nil
}

func calculation(revenue, expense, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(resultFileName, []byte(results), 0644)
}

// func writeResultsToFile(result float64) error {
	// resultText := fmt.Sprint(result)
	// if firstWrite {
		// firstWrite = false
		// return os.WriteFile("results.txt", []byte(resultText), 0644)
	// }
// 
	// f, err := os.OpenFile(resultFileName, os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
		// return err
	// }
	// defer f.Close()
// 
	// _, err = f.WriteString("\n" + resultText)
	// if err != nil {
		// return err
	// }
	// return nil
// 
// }