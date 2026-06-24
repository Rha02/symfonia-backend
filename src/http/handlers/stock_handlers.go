package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	DEFAULT_LIMIT = 50
	DEFAULT_SKIP  = 0
)

func (m *Repository) GetStocks(w http.ResponseWriter, r *http.Request) {
	limit, err := getQueryIntParamOrDefault(r, "limit", DEFAULT_LIMIT)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	skip, err := getQueryIntParamOrDefault(r, "skip", DEFAULT_SKIP)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := m.DB.GetStocks(limit, skip)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Failed to query database!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (m *Repository) SearchStock(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchKey := query.Get("value")

	res, err := m.DB.SearchStock(searchKey)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Failed to query database!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (m *Repository) GetStockBySymbol(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")

	stock, err := m.DB.GetStockBySymbol(symbol)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Failed to query database!")
		return
	}

	latestBar, err := m.Alpaca.GetStockLatest(symbol)
	if err != nil {
		log.Println(err)
		jsonError(w, http.StatusInternalServerError, "Failed to query Alpaca!")
		return
	}

	stock.LastPrice = latestBar.Bar.ClosingPrice

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func (m *Repository) GetStockTrend(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	if symbol == "" {
		jsonError(w, http.StatusBadRequest, "Missing request parameter: symbol")
		return
	}

	query := r.URL.Query()

	timeframe := query.Get("timeframe")
	if timeframe == "" {
		jsonError(w, http.StatusBadRequest, "Missing request parameter: timeframe")
		return
	}
	limitStr := query.Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Failed to parse request paramenter: limit")
		return
	}

	start := query.Get("start")
	if start == "" {
		jsonError(w, http.StatusBadRequest, "Missing request parameter: limit")
		return
	}

	data, err := m.Alpaca.GetStockTrend(symbol, timeframe, limit, start)
	if err != nil {
		log.Println(err)
		jsonError(w, http.StatusInternalServerError, "Failed to query Alpaca")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (m *Repository) GetPortfolioTrend(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	period := query.Get("period")
	if period == "" {
		jsonError(w, http.StatusBadRequest, "Missing request parameter: period")
		return
	}

	data, err := m.Alpaca.GetPortfolioTrend(period)
	if err != nil {
		log.Println(err)
		jsonError(w, http.StatusInternalServerError, "Failed to query Alpaca")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
