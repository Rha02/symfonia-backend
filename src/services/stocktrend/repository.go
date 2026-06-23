package stocktrend

import "github.com/Rha02/symfonia-backend/src/models"

type StockTrendRepository interface {
	GetStockLatest(symbol string) (*models.AlpacaLatestStockBar, error)
	GetStockTrend(symbol string, timeframe string, limit int, start string) (*[]models.AlpacaStockBar, error)
	GetPortfolioTrend(period string) (*models.AlpacaPortfolioResponse, error)
}
