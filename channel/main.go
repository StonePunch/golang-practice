package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Whenever a websites is checked, repeat the check after 2 seconds
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("The website:'", link, "'Might be down, Error:", err)
		c <- link
		return
	}

	fmt.Println("The website:'", link, "' is up")
	c <- link
}
