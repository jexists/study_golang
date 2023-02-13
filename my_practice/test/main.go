package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func main() {
	// Start a Selenium WebDriver instance.
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Printf("Failed to start selenium WebDriver: %s", err)
	}
	defer wd.Quit()

	// Load a page that contains an iframe.
	if err := wd.Get("https://cafe.naver.com/re4mo/1958074"); err != nil {
		fmt.Printf("Failed to load page: %s", err)
	}

	// Switch to the iframe.
	iframe, err := wd.FindElement(selenium.ByCSSSelector, "iframe")
	if err != nil {
		fmt.Printf("Failed to find iframe: %s", err)
	}
	wd.SwitchFrame(iframe)

	fmt.Println(iframe)
	// Extract information from the iframe using Goquery.
	// doc, err := goquery.NewDocumentFromReader(wd)
	// if err != nil {
	// 	fmt.Printf("Failed to create Goquery document: %s", err)
	// }

	// // Use Goquery to select and manipulate elements in the iframe.
	// doc.Find("p").Each(func(i int, s *goquery.Selection) {
	// 	text := s.Text()
	// 	fmt.Printf("Paragraph %d: %s\n", i+1, text)
	// })
}
