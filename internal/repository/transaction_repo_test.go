package repository

import (
	"encoding/json"
	"testing"
)

func TestTransactionRepo_ReadAll_EmptyTableMarshalsToArray(t *testing.T) {
	resetDB(t)

	repo := NewTransactionRepo(testDB)

	got, err := repo.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll() error: %v", err)
	}

	if len(got) != 0 {
		t.Fatalf("ReadAll() expected empty slice, got %d rows", len(got))
	}

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v", err)
	}

	if string(b) != "[]" {
		t.Errorf("json.Marshal() expected '[]', got '%s'", b)
	}
}
