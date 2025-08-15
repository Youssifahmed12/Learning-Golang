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
		"http://youtube.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println(l, err)
		c <- l
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println(l, "something wrong happened")
		c <- l
		return
	}

	fmt.Println(l, "is available")
	c <- l

}
