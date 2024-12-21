package main

import (
	"fmt"
	"net/http"
	"time"
)

func verifyLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("error: ", err)
		ch <- link
		return
	}

	fmt.Println(link, "is up!")
	ch <- link
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://apple.com",
	}

	ch := make(chan string)

	// start a goroutine for each link
	for _, link := range links {
		go verifyLink(link, ch)
	}

	// wait for all goroutines to finish
	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			verifyLink(link, ch)
		}(l)
	}
}
