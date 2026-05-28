package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rha02/symfonia-backend/src/dbrepo"
)

type Repository struct {
	DB dbrepo.DatabaseRepository
}

var Repo *Repository

func NewRepository(db dbrepo.DatabaseRepository) *Repository {
	return &Repository{
		db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func jsonError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func getQueryIntParamOrDefault(r *http.Request, key string, defaultVal int) (int, error) {
	query := r.URL.Query()
	str := query.Get(key)
	if str == "" {
		return defaultVal, nil
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultVal, fmt.Errorf("failed to parse url query param: '%s'", key)
	}
	return val, nil
}
