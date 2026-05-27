package models

type Stock struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
