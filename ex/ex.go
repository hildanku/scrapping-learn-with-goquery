package ex

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// func Check(url string) {
// 	// Request the HTML page.
// 	if url == "" {
// }

func Scrape() {
	res, err := http.Get("http://www.scrapethissite.com/pages/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("html body div#pages section div.container div.row div.col-md-6.col-md-offset-3 div.page h3.page-title").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Title %d: %s\n", i, title)
	})
}
