package sites

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FozzyScrape(u string) {
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
	text := doc.Find("span.delivery-information").Text()
	fmt.Println(text)
	if text != "Немає в наявності" {
		text = doc.Find(`[property="product:price:amount"]`).AttrOr("content", "Can not get a price")
		fmt.Println(text)
	}
}
