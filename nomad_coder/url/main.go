package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	errRequestFailed = errors.New("Request Failed")
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://github.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://okky.kr/",
		"https://www.airbnb.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

	// // var results = map[string]string // ERROR (초기화안되서) map == nil

	// // 초기화 시키기
	// // var results = map[string]string{}
	// var results = make(map[string]string)
	// urls := []string{
	// 	"https://github.com/",
	// 	"https://www.naver.com/",
	// 	"https://www.daum.net/",
	// 	"https://www.google.com/",
	// 	"https://www.facebook.com/",
	// 	"https://okky.kr/",
	// 	"https://www.airbnb.com/",
	// }

	// // results["hello"] = "hello"
	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }
	// // fmt.Println(results)

	// Top to Down
	// go sexyCount("joy")
	// sexyCount2("jexists")
}

// 보내는것만 가능
func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking ", url)
	// fmt.Println(<-c)
	// c <- result{}
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
		// fmt.Println(err, resp.StatusCode)
		// c <- result{url: url, status: "failed"}
	}
	c <- requestResult{url: url, status: status}

}

//
// func hitURL(url string) error {
// 	fmt.Println("Checking ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }

func sexyCount2(person string) {
	for i := 0; i < 5; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
