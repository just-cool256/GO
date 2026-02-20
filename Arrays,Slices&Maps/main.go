package main

import "fmt"

func main() {
	// We could make Go aware of the fact that we will add those two elements eventually. Because when we do, Go is able to, behind the scenes, create a bigger array right from the start, and it then doesn't have to recreate those arrays all the time, which is a bit more efficient. And you can tell Go that you need a bigger array behind the scenes by using the make function.
	// userNames := []string{}
	userNames := make([]string, 2, 5) //Go then behind the scenes, create an array with length 2 and capacity 5, where both elements will be initially null, empty slots, but where you can now assign elements to those slots
	// The capacity is essentially the maximum number of array items, and therefore controls how much memory space will be allocated behind the scenes by Go for this array.
	// So this make function call will make sure that Go allocates enough memory space for a string array that takes five items, and it will create an array with 2 empty slots.

	userNames[0] = "Julie"
	userNames[1] = "John"
	userNames = append(userNames, "Max", "Manuel")
	// But now with pre allocated memory, appending items here is more efficient, because Go now doesn't have to go to the memory and allocate new space, instead it can use the existing space because we already reserved enough space. And only once we go beyond 5 capacity limit here, it'll have to allocate new space.

	fmt.Println(userNames)

	// You can use make for making slices and maps both
	// courseRatings := map[string]float64{}
	// courseRatings["go"] = 4.7
	// courseRatings["react"] = 4.8

	// with map we can only pass 1 additional argument in make
	// bcz here we can't set any empty slots. Instead we can justify the intended length of that map
	// So, Go can go ahead and pre-allocate memory.
	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["ansible"] = 4.7
	// courseRatings["angular"] = 4.8 // go would have to reallocate memory only when

	fmt.Println(courseRatings)

}
