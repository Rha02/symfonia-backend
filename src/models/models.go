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

type AlpacaPortfolioResponse struct {
	Timestamp     []int     `json:"timestamp"`
	Equity        []float32 `json:"equity"`
	ProfitLoss    []float32 `json:"profit_loss"`
	ProfitLossPct []float32 `json:"profit_loss_pct"`
	BaseValue     float32   `json:"base_value"`
}

type AlpacaLatestStockBar struct {
	Symbol string         `json:"symbol"`
	Bar    AlpacaStockBar `json:"bar"`
}
