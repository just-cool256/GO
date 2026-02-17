package main

import "fmt"

func main() {
	// Map is always dynamic
	website := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}

	fmt.Println(website)           // map[Amazon Web Services:https://aws.amazon.com Google:https://www.google.com]
	fmt.Println(website["Google"]) // https://www.google.com

	website["Facebook"] = "https://www.facebook.com"
	fmt.Println(website) // map[Amazon Web Services:https://aws.amazon.com Facebook:https://www.facebook.com Google:https://www.google.com]

	delete(website, "Google")
	fmt.Println(website) // map[Amazon Web Services:https://aws.amazon.com Facebook:https://www.facebook.com]
}

/*
difference between map and struct
1. For maps, you can use anything as a key like string, int, float or even some array or a struct as a key
2. With structs, you have predefined data structures, we can't add or delete fields afterwards
You don't use structs to manage multiple values of the same kind with different keys, but instead you use structs to describe data entities in your programs.
And a map on the other hand is used if you have a collection of values, which probably describe the same thing, though technically, not necessarily, but then you can assign your own labels, your own keys, and use any kinds of values for those keys, that make sense to you
*/