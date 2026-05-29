package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rha02/symfonia-backend/src/models"
)

const timeout = 5 * time.Second

type postgresdbRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) DatabaseRepository {
	return &postgresdbRepo{
		db: db,
	}
}

// GetStocks implements [DatabaseRepository].
func (m *postgresdbRepo) GetStocks(limit int, skip int) (*[]models.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var stocks []models.Stock

	stmt := `
		SELECT id, symbol, name, created_at FROM stock
		LIMIT $1 OFFSET $2;
	`

	rows, err := m.db.QueryContext(ctx, stmt, limit, skip)
	if err != nil {
		return &stocks, nil
	}

	for rows.Next() {
		var stock models.Stock

		if err := rows.Scan(&stock.ID, &stock.Symbol, &stock.Name, &stock.CreatedAt); err != nil {
			return &stocks, err
		}

		stocks = append(stocks, stock)
	}

	return &stocks, nil
}

// SearchStock implements [DatabaseRepository].
func (m *postgresdbRepo) SearchStock(searchKey string) (*[]models.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stocks := make([]models.Stock, 0)

	stmt := `
		SELECT id, symbol, name, created_at FROM stock
		WHERE symbol ILIKE $1
		OR name ILIKE $1
		ORDER BY
			CASE WHEN symbol ILIKE $1 THEN 0 ELSE 1 END
		LIMIT $2;
	`

	searchKeyWildcard := "%" + searchKey + "%"
	searchLimit := 5

	rows, err := m.db.QueryContext(ctx, stmt, searchKeyWildcard, searchLimit)
	if err != nil {
		return &stocks, err
	}

	for rows.Next() {
		var stock models.Stock

		if err := rows.Scan(&stock.ID, &stock.Symbol, &stock.Name, &stock.CreatedAt); err != nil {
			return &stocks, err
		}

		stocks = append(stocks, stock)
	}

	return &stocks, nil
}
