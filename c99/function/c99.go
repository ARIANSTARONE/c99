package function

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func C99(url string) {
	// Get the current date and subtract 7 days
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	// Construct the full URL with the calculated date
	var baseurl = "https://subdomainfinder.c99.nl/scans/" + sevenDaysAgo + "/" + url

	res, err := http.Get(baseurl)
	if err != nil {
		log.Fatal("Error fetching page:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("Failed to fetch page, status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body:", err)
	}

	ipRegex := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	doc.Find("div.well.mt-5").Each(func(index int, item *goquery.Selection) {
		item.Find("a").Each(func(indexA int, a *goquery.Selection) {
			text := a.Text()
			if !strings.EqualFold(text, "none") && !ipRegex.MatchString(text) {
				fmt.Println("", text)
			}
		})
	})
}
