package model

import "time"

type Transaction struct {
	ID       int       `json:"id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
	Type     string    `json:"type"`
}
