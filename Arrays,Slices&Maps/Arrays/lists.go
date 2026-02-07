package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	// productNames = [4]string {"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[2])
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3][1:]
	fmt.Println(featuredPrices)

	// featuredPrices = prices[-1:] - can't work with negative indicies
	// featuredPrices = prices[1:7] - can't pick higher end index than the original array has
}
