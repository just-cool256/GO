package main

func main() {
	numbers := []int{1, 10, 15}
	// sum := sumup(numbers)
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(numbers...)
	println("The sum of the numbers is:", sum)
	println("The sum of the numbers is:", anotherSum)
}

// func sumup(startingValue int, numbers ...int) int {}
func sumup(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
