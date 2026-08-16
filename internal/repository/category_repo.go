package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/filipbabicdev/finance-tracker-api/internal/model"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) Create(c *model.Category) error {
	query := "INSERT INTO categories (name, type, bucket) VALUES ($1, $2, $3) RETURNING id, created_at"
	if err := r.db.QueryRow(query, c.Name, c.Type, c.Bucket).Scan(&c.ID, &c.CreatedAt); err != nil {
		return fmt.Errorf("create category: %w", err)
	}
	return nil
}

func (r *CategoryRepo) GetAll() ([]model.Category, error) {
	query := "SELECT id, name, type, bucket, created_at FROM categories ORDER BY name"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("get all categories: %w", err)
	}
	defer rows.Close()

	categories := []model.Category{}
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Type, &c.Bucket, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate categories: %w", err)
	}

	return categories, nil
}

func (r *CategoryRepo) GetByID(id int) (*model.Category, error) {
	query := "SELECT id, name, type, bucket, created_at FROM categories WHERE id = $1"

	var c model.Category
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Type, &c.Bucket, &c.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrCategoryNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get category by id: %w", err)
	}

	return &c, nil
}
