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
}
