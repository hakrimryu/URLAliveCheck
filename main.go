package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	urls := []string{
		"https://google.com/",
		"https://naver.com/",
		"https://wrongaddress.com/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Alive! : ", url)

	response, err := http.Get(url)

	if err == nil || response.StatusCode >= 400 {
		fmt.Println("Dead! : ", url)
		return errRequestFailed
	}

	return nil
}
