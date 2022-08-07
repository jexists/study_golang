package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://cafe.naver.com/re4mo?iframe_url=/ArticleList.nhn%3Fsearch.clubid=13988016%26search.menuid=330%26search.boardtype=L"

func main() {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	search := doc.Find("body")
	search.Each(func(i int, s *goquery.Selection) {
		// id = getArticleID(s)
		fmt.Println(s.Html())
	})
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Fail Status: ", res.StatusCode)
	}
}
