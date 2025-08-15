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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(l string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Println(l, err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println(l, "something wrong happened")
		return
	}

	fmt.Println(l, "is available")

}
