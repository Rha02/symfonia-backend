package stocktrend

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/Rha02/symfonia-backend/src/models"
)

type alpacaRepo struct {
	apiKey    string
	apiSecret string
}

const alpacaURL = "https://data.alpaca.markets/v2"

func NewAlpacaRepo(apiKey string, apiSecret string) StockTrendRepository {
	return &alpacaRepo{apiKey: apiKey, apiSecret: apiSecret}
}

// GetStockTrend implements [StockTrendRepository].
func (a *alpacaRepo) GetStockTrend(symbol string, timeframe string, limit int, start string) (*[]models.AlpacaStockBar, error) {
	uri := alpacaURL + "/stocks/bars"

	res, err := http.Get(uri)
	if err != nil {
		log.Println("Error querying alpaca")
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body")
		return nil, err
	}

	var resBody models.AlpacaBarsResponse

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		log.Println("Error unmarshalling json")
		return nil, err
	}

	data, ok := resBody.Bars[symbol]
	if !ok {
		log.Printf("Error, no trend data for ")
		return nil, errors.New("missing data for symbol")
	}

	return &data, nil
}
