package processor

import (
	"time"
)

type PriceData struct {
	Currency string
	Price    float64
}

type ProcessedData struct {
	Currency  string
	Price     float64
	Timestamp time.Time
}

func ProcessData(priceCh <-chan PriceData, processedDataCh chan<- ProcessedData) {

	accumulatedPrices := make(map[string]struct {
		Total float64
		Count int
	})
	for priceData := range priceCh {
		accumulatedPrices[priceData.Currency] = struct {
			Total float64
			Count int
		}{
			Total: accumulatedPrices[priceData.Currency].Total + priceData.Price,
			Count: accumulatedPrices[priceData.Currency].Count + 1,
		}

		averagePrice := accumulatedPrices[priceData.Currency].Total / float64(accumulatedPrices[priceData.Currency].Count)

		processedData := ProcessedData{
			Currency:  priceData.Currency,
			Price:     averagePrice,
			Timestamp: time.Now(),
		}

		//log.Println(processedData)

		processedDataCh <- processedData
	}
}
