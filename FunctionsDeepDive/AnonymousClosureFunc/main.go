package main

import "fmt"

// When you have a case where a function wants another function as a parameter value or where a function returns a function, then you often can save some effort and time by using anonymous functions, which is a feature that allows you to define a function just in time when you need it instead of in advance.

// Closures use anonymous function and use a specific aspect of anonymous function
func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func (number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// functions that produces other functions with different configurations.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
// Every anonymous function is a closure - it means if you use a variable from a scope in which the function is created(from outer scope), then the value of that outer scope variable or parameter is locked in into this created function