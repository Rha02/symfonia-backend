package models

type Stock struct {
	ID        int    `json:"id"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
