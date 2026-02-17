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
