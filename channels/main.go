package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://github.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://google.com",
		"http://apple.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	for l := range c {
		// fmt.Println(<-c)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(time.Second * 5)
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " must be down")
		c <- link
		return
	}

	fmt.Println(link, " is up")
	c <- link
}
