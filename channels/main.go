package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string, c chan string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println(l, err)
		c <- "Error Occured"
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println(l, "something wrong happened")
		c <- "definitely something wrong"
		return
	}

	fmt.Println(l, "is available")
	c <- "all perfect"

}
