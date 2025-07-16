package sites

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func EpicentrkScrape(u string) {
	res, err := http.Get(u)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Request status: %s", res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := doc.Find("[class='_A7y+idsw']").Text()
	fmt.Println(text)
	text = doc.Find("data[itemprop=\"price\"]").Last().AttrOr("content", "Cann't get a price")
	fmt.Println(text)
}
