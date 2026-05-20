package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID       int             `json:"id"`
	Amount   decimal.Decimal `json:"amount"`
	Category string          `json:"category"`
	Date     time.Time       `json:"date"`
	Type     string          `json:"type"`
}
