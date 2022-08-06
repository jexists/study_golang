package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	id       string
	location string
	title    string
	link     string
}

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	// fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	// fmt.Println(len(jobs))
	fmt.Println("끝! ", len(jobs))
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	checkErr(err)

	// if res.StatusCode != 200 {
	// 	log.Fatalln("Request Fail Status: ", res.StatusCode)
	// }
	checkCode(res)

	defer res.Body.Close()
	// 메모리 새어나가는거 막기

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html()) // html
		// fmt.Println(s.Find("a")) // pointer
		// fmt.Println(s.Find("a").Html()) // 첫번째 HTML만 나옴
		// fmt.Println(s.Find("a").Length()) // 숫자
		pages = s.Find("a").Length()
	})

	return pages
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	// fmt.Println(pageURL)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		// fmt.Println(job)
		jobs = append(jobs, job)
	})

	fmt.Println(len(jobs))
	return jobs
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Find(".jobTitle a").Attr("data-jk")
	title := cleanString(card.Find(".jobTitle a span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	link := "https://kr.indeed.com/viewjob?jk=" + id
	// fmt.Println(id, title, location)
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		link:     link}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Link"}
	err = w.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.link}
		err := w.Write(jobSlice)
		checkErr(err)

	}
}
