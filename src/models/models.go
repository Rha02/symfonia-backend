package models

type Stock struct {
	ID        int    `json:"id"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type AlpacaBarsResponse struct {
	Bars map[string][]AlpacaStockBar `json:"bars"`
}

type AlpacaStockBar struct {
	Timestamp    string  `json:"t"`
	OpeningPrice float32 `json:"o"`
	HighPrice    float32 `json:"h"`
	LowPrice     float32 `json:"l"`
	ClosingPrice float32 `json:"c"`
	Volume       int     `json:"v"`
	TradeCount   int     `json:"n"`
	AveragePrice float32 `json:"vw"`
}
