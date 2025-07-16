package internal

type AuchanJson struct {
	Data struct {
		Products struct {
			Items []struct {
				StockStatus string `json:"stock_status"`
				PriceRange  struct {
					MinimumPrice struct {
						FinalPrice struct {
							Value float64 `json:"value"`
						} `json:"final_price"`
					} `json:"minimum_price"`
				} `json:"price_range"`
			} `json:"items"`
		} `json:"products"`
	} `json:"data"`
}
