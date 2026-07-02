package handler

import (
	"time"

	"github.com/filipbabicdev/finance-tracker-api/internal/model"
	"github.com/shopspring/decimal"
)

type TransactionRequest struct {
	Amount      decimal.Decimal `json:"amount" binding:"required"`
	Category    string          `json:"category" binding:"required"`
	Date        time.Time       `json:"date" binding:"required"`
	Type        string          `json:"type" binding:"required,oneof=income expense"`
	Description string          `json:"description"`
	Merchant    *string         `json:"merchant"`
	Currency    string          `json:"currency" binding:"omitempty,iso4217"`
}

func (r *TransactionRequest) ToModel() *model.Transaction {
	currency := r.Currency
	if currency == "" {
		currency = "RSD"
	}

	return &model.Transaction{
		Amount:      r.Amount,
		Category:    r.Category,
		Date:        r.Date,
		Type:        r.Type,
		Description: r.Description,
		Merchant:    r.Merchant,
		Currency:    currency,
	}
}
