package repository

import (
	"database/sql"

	"github.com/filipbabicdev/finance-tracker-api/internal/model"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) Create(t *model.Transaction) error {
	query := "INSERT INTO transactions (amount, category, date, type) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.db.QueryRow(query, t.Amount, t.Category, t.Date, t.Type).Scan(&t.ID)
}

func (r *TransactionRepo) ReadAll() ([]model.Transaction, error) {
	query := "SELECT id, amount, category, date, type FROM transactions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.ID, &t.Amount, &t.Category, &t.Date, &t.Type); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepo) Update(id int, t *model.Transaction) error {
	query := "UPDATE transactions SET amount=$1, category=$2, date=$3, type=$4 WHERE id=$5"
	_, err := r.db.Exec(query, t.Amount, t.Category, t.Date, t.Type, id)
	return err
}

func (r *TransactionRepo) Delete(id int) error {
	query := "DELETE FROM transactions WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}
