package main

import (
	"fmt"
	"net/http"
)

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

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
