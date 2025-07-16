package sites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"parser/internal"
)

const apiURL = `https://express.auchan.ua/graphql/`

func AuchanScrape(URLKey string) {
	reqJson := fmt.Sprintf(`{
    "operationName": "getProductDetail",
    "variables": {
      "urlKey": "%s"
    },
    "query": "query getProductDetail($urlKey: String) { products(filter: {url_key: {eq: $urlKey}}) { items { stock_status price_range { minimum_price { final_price { value } } } } } }"
  }`, URLKey)
	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var resJson internal.AuchanJson
	err = json.Unmarshal(body, &resJson)
	if err != nil {
		log.Fatalln(err)
	}

	if resJson.Data.Products.Items[0].StockStatus == "IN_STOCK" {
		fmt.Println("В наявності")
	} else {
		fmt.Println("Немає в наявності")
	}
	fmt.Println(resJson.Data.Products.Items[0].PriceRange.MinimumPrice.FinalPrice.Value)
}
