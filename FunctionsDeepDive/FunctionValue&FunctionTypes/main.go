package main

import (
	"fmt"
)

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

// Functions are first class values in go
// We can use functions themselves as parameter values for other functions
func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := make([]int, len(*numbers))
	for i, val := range *numbers {
		dNumbers[i] = transform(val)
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
