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
	query := "INSERT INTO transactions (amount, category, date, type) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, t.Amount, t.Category, t.Date, t.Type)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = int(id)
	return nil
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
	query := "UPDATE transactions SET amount=?, category=?, date=?, type=? WHERE id=?"
	_, err := r.db.Exec(query, t.Amount, t.Category, t.Date, t.Type, id)
	return err
}

func (r *TransactionRepo) Delete(id int) error {
	query := "DELETE FROM transactions WHERE id=?"
	_, err := r.db.Exec(query, id)
	return err
}
