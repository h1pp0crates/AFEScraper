package sites

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func EpicentrkScrape(u string) string {
	res, err := http.Get(u)
	if err != nil {
		log.Println("EpicentrkScrape:\nGet request error:\n", err)
		return "Error"
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("EpicentrkScrape:\nRequest status: %s", res.Status)
		return "Error"
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("EpicentrkScrape:\nres.Body error:\n", err)
		return "Error"
	}

	if doc.Find("[class='_A7y+idsw']").Text() == "В наявності" {
		return doc.Find("data[itemprop=\"price\"]").Last().AttrOr("content", "Cann't get a price") + " грн"
	}

	return "Немає в наявності"
}
