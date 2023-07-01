package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PriceData struct {
	Currency string
	Price    float64
}

type priceResponse struct {
	USD float64 `json:"usd"`
}

func TrackPrices(priceCh chan<- PriceData) {
	currencies := []string{"bitcoin", "ethereum", "ripple", "litecoin", "bitcoin-cash"}

	client := http.Client{
		Timeout: 7 * time.Second,
	}

	for {
		for _, currency := range currencies {
			url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", currency)

			resp, err := client.Get(url)

			if err != nil {
				fmt.Printf("Error fetching price for %s: %v", currency, err)
				continue
			}

			var result map[string]priceResponse

			err = json.NewDecoder(resp.Body).Decode(&result)

			if err != nil {
				fmt.Printf("Error decoding price data for %s: %v", currency, err)
				continue
			}

			price := result[currency].USD

			/*price, err := strconv.ParseFloat(pricStr, 64)

			if err != nil {
				log.Printf("Error prasing price for %s: %v\n", currency, err)
				continue
			} */

			resp.Body.Close()

			priceCh <- PriceData{
				Currency: currency,
				Price:    price,
			}
		}

		time.Sleep(1 * time.Second)
	}
}
