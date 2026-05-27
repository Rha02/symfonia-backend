package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rha02/symfonia-backend/src/models"
)

func (m *Repository) GetStocks(w http.ResponseWriter, r *http.Request) {
	res := []models.Stock{
		{
			ID: 1, Symbol: "STCK1", Name: "Stock Test 1",
		},
		{
			ID: 2, Symbol: "STCK2", Name: "Stock Test 2",
		},
		{
			ID: 3, Symbol: "STIX", Name: "Dummy Stock Test",
		},
		{
			ID: 4, Symbol: "TUSK", Name: "My Dummy Stock",
		},
		{
			ID: 5, Symbol: "POS", Name: "Test Stock",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
