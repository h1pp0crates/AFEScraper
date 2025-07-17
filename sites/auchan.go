package sites

import (
	"AFEScraper/internal"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func AuchanScrape(URLKey string) string {
	apiURL := "https://express.auchan.ua/graphql/" // Replace with actual API URL

	reqJson := fmt.Sprintf(`{
    "operationName": "getProductDetail",
    "variables": {
      "urlKey": "%s"
    },
    "query": "query getProductDetail($urlKey: String) { products(filter: {url_key: {eq: $urlKey}}) { items { stock_status price_range { minimum_price { final_price { value } } } } } }"
  }`, URLKey)

	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		log.Println("AuchanScrape\nPost request error:\n", err)
		return "Error"
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("AuchanScrape:\nres.Body error:\n", err)
		return "Error"
	}

	var resJson internal.AuchanJson
	err = json.Unmarshal(body, &resJson)
	if err != nil {
		log.Println("AuchanScrape:\njson.Unmarshal error:\n", err)
		return "Error"
	}

	if len(resJson.Data.Products.Items) == 0 {
		return "Помилка, немає даних в resJson"
	}

	if resJson.Data.Products.Items[0].StockStatus == "IN_STOCK" {
		return fmt.Sprintf("%.2f грн", resJson.Data.Products.Items[0].PriceRange.MinimumPrice.FinalPrice.Value)
	}
	return "Немає в наявності"
}
