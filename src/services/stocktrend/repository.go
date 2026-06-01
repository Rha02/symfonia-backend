package stocktrend

import "github.com/Rha02/symfonia-backend/src/models"

type StockTrendRepository interface {
	GetStockTrend(symbol string, timeframe string, limit int, start string) (*[]models.AlpacaStockBar, error)
}
