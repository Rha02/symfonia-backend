package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	DEFAULT_LIMIT = 50
	DEFAULT_SKIP  = 0
)

func (m *Repository) GetStocks(w http.ResponseWriter, r *http.Request) {
	limit, err := getQueryIntParamOrDefault(r, "limit", DEFAULT_LIMIT)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
	}

	skip, err := getQueryIntParamOrDefault(r, "skip", DEFAULT_SKIP)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
	}

	res, err := m.DB.GetStocks(limit, skip)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Failed to query database!")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
