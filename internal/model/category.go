package model

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Bucket    *string   `json:"bucket"`
	CreatedAt time.Time `json:"created_at"`
}
