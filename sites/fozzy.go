package sites

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FozzyScrape(u string) string {
	res, err := http.Get(u)
	if err != nil {
		log.Println("FozzyScrape:\nGet request error:\n", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("FozzyScrape:\nRequest status: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("FozzyScrape:\nres.Body error:\n", err)
	}

	if doc.Find("span.delivery-information").Text() != "Немає в наявності" {
		return doc.Find(`[property="product:price:amount"]`).AttrOr("content", "Can not get a price") + " грн"
	}
	return "Немає в наявності"
}
