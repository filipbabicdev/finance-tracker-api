package repository

import (
	"errors"
	"testing"

	"github.com/filipbabicdev/finance-tracker-api/internal/model"
)

func TestCategoryRepo_Create(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	c := model.Category{Name: "Groceries", Type: "expense", Bucket: strPtr("needs")}
	if err := repo.Create(&c); err != nil {
		t.Fatalf("Create() error: %v", err)
	}

	if c.ID == 0 {
		t.Error("Create() did not populate ID field")
	}

	if c.CreatedAt.IsZero() {
		t.Error("Create() did not populate CreatedAt field")
	}
}

func TestCategoryRepo_Create_RejectsIncomeWithBucket(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	c := model.Category{Name: "Salary", Type: "income", Bucket: strPtr("needs")}
	if err := repo.Create(&c); err == nil {
		t.Fatal("Create() accepted income category with a bucket, want constraint violation")
	}
}

func TestCategoryRepo_GetAll_EmptyReturnsEmptySlice(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	got, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error: %v", err)
	}
	if got == nil {
		t.Error("GetAll() returned nil slice, want empty slice")
	}
	if len(got) != 0 {
		t.Errorf("GetAll() returned %d categories, want 0", len(got))
	}
}

func TestCategoryRepo_GetAll_OrdersByName(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	for _, name := range []string{"Utilities", "Groceries", "Entertainment"} {
		c := model.Category{Name: name, Type: "expense", Bucket: strPtr("needs")}
		if err := repo.Create(&c); err != nil {
			t.Fatalf("Create() seed %q error: %v", name, err)
		}
	}

	got, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error: %v", err)
	}

	want := []string{"Entertainment", "Groceries", "Utilities"}
	if len(got) != len(want) {
		t.Fatalf("GetAll() returned %d categories, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i].Name != want[i] {
			t.Errorf("GetAll() index %d: got %q, want %q", i, got[i].Name, want[i])
		}
	}
}

func TestCategoryRepo_GetByID(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	created := model.Category{Name: "Groceries", Type: "expense", Bucket: strPtr("needs")}
	if err := repo.Create(&created); err != nil {
		t.Fatalf("Create() error: %v", err)
	}

	got, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID() error: %v", err)
	}

	if got.Name != created.Name {
		t.Errorf("GetByID() got Name %q, want %q", got.Name, created.Name)
	}
}

func TestCategoryRepo_GetByID_NotFound(t *testing.T) {
	resetDB(t)
	repo := NewCategoryRepo(testDB)

	_, err := repo.GetByID(9999)
	if !errors.Is(err, ErrCategoryNotFound) {
		t.Fatalf("GetByID() error: got %v, want ErrNotFound", err)
	}
}

func strPtr(s string) *string {
	return &s
}
