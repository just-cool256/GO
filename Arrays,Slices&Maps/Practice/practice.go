package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"Trekking", "Traveling", "Coding"}
	fmt.Println(hobbies)
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("Second and third hobby:", hobbies[1:3])

	// Creating slices in two different ways
	slice1 := hobbies[0:2] // Slice containing the first and second element
	slice2 := hobbies[:2]  // Another way to create the same slice

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	// Re-slicing to contain the second and last element
	slice1 = slice1[1:3] // Now slice1 contains the second and last element of the original array
	fmt.Println("Re-sliced Slice 1:", slice1)

	// Dynamic array (slice) for course goals
	courseGoals := []string{"Learn Go", "Build a project"}
	fmt.Println("Course Goals:", courseGoals)

	// Updating the second goal and adding a third goal
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Contribute to open source")
	fmt.Println("Updated Course Goals:", courseGoals)

	// Creating a dynamic list of products
	products := []Product{
		{"Laptop", 1, 999.99},
		{"Smartphone", 2, 499.99},
	}
	fmt.Println("Products:", products)

	products = append(products, Product{"Tablet", 3, 299.99})
	fmt.Println("Updated Products:", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
