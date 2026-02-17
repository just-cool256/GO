package main

import "fmt"

func main() {
	website := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}

	fmt.Println(website) // map[Amazon Web Services:https://aws.amazon.com Google:https://www.google.com]
}
