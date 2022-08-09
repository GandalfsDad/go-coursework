package main

import "fmt"

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.org",
	}

	for _, website := range websites {
		fmt.Println(website)
	}
}
