package main

import (
	"fmt"
)

// Function types
type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

// Functions are first class values in go
// We can use functions themselves as parameter values for other functions
// Function can also return functions as return values
func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers1 := transformNumbers(&numbers, transformFn1)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformFn2)
	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := make([]int, len(*numbers))
	for i, val := range *numbers {
		dNumbers[i] = transform(val)
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
