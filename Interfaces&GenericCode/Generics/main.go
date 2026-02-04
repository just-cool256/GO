package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
	// result + 1 - invalid operation when func add(a, b interface{}) interface{} {}
}

// func add[T interface{}](a, b T) T{}
func add[T int | float64 | string ](a, b T) T {
	return a + b
}

// func add(a, b interface{}) interface{} {
	//interface{} - this type is a bit too wide. Accepting any kind of value is a bit too wide, and in addition another problem we have here is that the return type also is too unspecific
	// return a + b - will not work
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)
// 
	// if aIsInt && bIsInt {
		// return aInt + bInt
	// }
// 
	// aFloat, aIsFloat := a.(int)
	// bFloat, bIsFloat := b.(int)
// 
	// if aIsFloat && bIsFloat {
		// return aFloat + bFloat
	// }
// 
	// aString, aIsString := a.(int)
	// bString, bIsString := b.(int)
// 
	// if aIsString && bIsString {
		// return aString + bString
	// }
	// return nil
// }