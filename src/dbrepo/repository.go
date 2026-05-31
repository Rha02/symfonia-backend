package dbrepo

import "github.com/Rha02/symfonia-backend/src/models"

type DatabaseRepository interface {
	GetStocks(limit int, skip int) (*[]models.Stock, error)
	SearchStock(searchKey string) (*[]models.Stock, error)
	GetStockBySymbol(symbol string) (*models.Stock, error)
}
