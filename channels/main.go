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

	c := make(chan string)

	for _, website := range websites {
		go checkLink(website, c)
	}

	v := []int{1, 2, 3, 4, 5}

	for i := range v {
		fmt.Println(<-c)
		fmt.Println(i)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
