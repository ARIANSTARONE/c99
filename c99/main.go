package main

import (
	"flag"
	"log"
	"main.go/function"
	_ "main.go/function"
)

func main() {

	urlFlag := flag.String("u", "", "URL to scrape")
	flag.Parse()

	if *urlFlag == "" {
		log.Fatal("URL flag is required")
	}
	function.C99(*urlFlag)
}

// ScrapeWebsite performs the web scraping and prints results
//func ScrapeWebsite(url string) {
//	var baseurl = "https://subdomainfinder.c99.nl/scans/2024-04-23/" + url
//	res, err := http.Get(baseurl)
//	if err != nil {
//		log.Fatal("Error fetching page:", err)
//	}
//	defer res.Body.Close()
//
//	if res.StatusCode != 200 {
//		log.Fatalf("Failed to fetch page, status code: %d", res.StatusCode)
//	}
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal("Error loading HTTP response body:", err)
//	}
//
//	ipRegex := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
//
//	doc.Find("div.well.mt-5").Each(func(index int, item *goquery.Selection) {
//		item.Find("a").Each(func(indexA int, a *goquery.Selection) {
//			text := a.Text()
//			if !strings.EqualFold(text, "none") && !ipRegex.MatchString(text) {
//				fmt.Println("", text)
//			}
//		})
//	})
//}
