package main

import "fmt"

// With Slices you can create dynamic arrays
func main() {
	// prices := []float64{} //empty slice
	// In this manner Go will automtically create a slice for you, and since a slice is based on an array, it will also create an array for you behind the scenes, but it will automtically ditch that array and create a new array if your slice grows beyond the balance of that behind the scenes stored array.

	prices := []float64{10.99, 8.99} // In this way we can create a slice without having to specify the length of it, and we can also add as many values as we want in there.
	fmt.Println(prices[0:1])
}

/*
func main() {
	var productNames [4]string = [4]string{"A Book"}
	// productNames = [4]string {"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0} // Downside of using that array was that we had to define ahead of time how many values we'll have in there.
	fmt.Println(prices)
	fmt.Println(prices[2])
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	fmt.Println()

	fmt.Println("senerio - 1")
	fmt.Println("featuredPrices = prices[1:3]")
	featuredPrices = prices[1:3]
	// Slices are like a reference, like a window into an array. When we created an array like prices that array is stored in memory.
	// When we then create a slice based on that array, we get a window into that array.
	// Therefore if you would modify an element in a slice, we would also modify the same element in the original array.

	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices)
	// fmt.Println(prices)
	// Important Implication: When you create a slice, you don't copy the original array, so you don't have that copy in memory that occupies extra memory space.
	// Instead you only have 1 array in memory and your slice is just a tiny reference to a part of that array.
	// Hence it's a very memory efficient way of selecting parts of an array or editing parts of an array.

	// Go also saves some metadata for our slices that can be useful to look into.
	// For every slice we got a length and a capacity

	fmt.Println(len(featuredPrices), cap(featuredPrices)) //len - 2, cap - 3
	// length - no. of items in a slice or array
	// cap - tells us - Go memorized that there is more content available to the right of the selected slice.
	// we can reslice to pick up more elements from right but that is not possible from left

	// fmt.Println()
	// fmt.Println("featuredPrices = featuredPrices[1:]")
	// featuredPrices = featuredPrices[1:]
	// fmt.Println(featuredPrices)
	// fmt.Println(len(featuredPrices), cap(featuredPrices))

	fmt.Println()
	fmt.Println("featuredPrices = featuredPrices[:3]")
	featuredPrices = featuredPrices[:3] //- gave panic error
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	fmt.Println("senerio-2")
	featuredPrices = prices[1:]             //{199.99, 45.99, 20.0}
	highlightedPrices := featuredPrices[:1] //{199.99}
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices) // {199.99 45.99 20}
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// featuredPrices = prices[:]
	// fmt.Println(featuredPrices)

	// featuredPrices = prices[:3][1:]
	// fmt.Println(featuredPrices)

	// featuredPrices = prices[-1:] - can't work with negative indicies
	// featuredPrices = prices[1:7] - can't pick higher end index than the original array has
}

// When we suddenly select more than was previously stored in a slice - that works
// Bcz if you select more to the end of an array, you can do that bcz go memorizes that, but it only work for more values to the right of the existing slice
*/