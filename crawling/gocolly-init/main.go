package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// go get -u github.com/gocolly/colly
	// http://go-colly.org/docs/introduction/install/

	c := colly.NewCollector(
		colly.AllowedDomains("https://www4.kitbag.com/en/soccer/england-national-team/england-away-stadium-shirt-2022/o-3623+t-81616060+p-01363635655+z-8-3887643823?_ref=p-PDP:m-YMAL:pi-PDP_RECOMMENDATIONS_2:i-r0c2:po-2"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
