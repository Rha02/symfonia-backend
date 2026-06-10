package stocktrend

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Rha02/symfonia-backend/src/models"
)

type alpacaRepo struct {
	apiKey    string
	apiSecret string
}

const marketURL = "https://data.alpaca.markets/v2"
const apiURL = "https://paper-api.alpaca.markets/v2"

const timeout = 10 * time.Second

func NewAlpacaRepo(apiKey string, apiSecret string) StockTrendRepository {
	return &alpacaRepo{apiKey: apiKey, apiSecret: apiSecret}
}

// GetStockTrend implements [StockTrendRepository].
func (a *alpacaRepo) GetStockTrend(symbol string, timeframe string, limit int, start string) (*[]models.AlpacaStockBar, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	uri, _ := url.Parse(marketURL + "/stocks/bars")

	limitStr := strconv.Itoa(limit)

	params := url.Values{}
	params.Add("symbols", symbol)
	params.Add("timeframe", timeframe)
	params.Add("limit", limitStr)
	params.Add("start", start)

	uri.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		log.Println("Error creating request")
		return nil, err
	}
	req.Header.Set("apca-api-key-id", a.apiKey)
	req.Header.Set("apca-api-secret-key", a.apiSecret)

	cli := &http.Client{}

	res, err := cli.Do(req)
	if err != nil {
		log.Println("Error querying Alpaca", err)
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

// GetPortfolioTrend implements [StockTrendRepository].
func (a *alpacaRepo) GetPortfolioTrend(period string) (*models.AlpacaPortfolioResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	uri, _ := url.Parse(apiURL + "/account/portfolio/history")

	params := url.Values{}
	params.Add("period", period)

	uri.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		log.Println("Error creating request")
		return nil, err
	}
	req.Header.Set("apca-api-key-id", a.apiKey)
	req.Header.Set("apca-api-secret-key", a.apiSecret)

	cli := &http.Client{}

	res, err := cli.Do(req)
	if err != nil {
		log.Println("Error querying Alpaca", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body")
		return nil, err
	}

	var resBody *models.AlpacaPortfolioResponse

	err = json.Unmarshal(body, &resBody)
	if err != nil {
		log.Println("Error unmarshalling json")
		return nil, err
	}

	return resBody, nil
}
