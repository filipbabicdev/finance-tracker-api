package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewTransactionHandler(nil)
	SetupRoutes(r, h)
	return r
}

func postJSON(t *testing.T, r *gin.Engine, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/transactions", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateRejectsInvalidType(t *testing.T) {
	r := setupRouter()
	w := postJSON(t, r, `{"amount":"100.00","category":"food","date":"2026-07-02T10:00:00Z","type":"foo"}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d, body: %s", w.Code, w.Body.String())
	}
}

func TestCreateRejectsInvalidCurrency(t *testing.T) {
	r := setupRouter()
	w := postJSON(t, r, `{"amount":"100.00","category":"food","date":"2026-07-02T10:00:00Z","type":"expense","currency":"xyz"}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d, body: %s", w.Code, w.Body.String())
	}
}

func TestToModelDefaultsCurrencyToRSD(t *testing.T) {
	req := TransactionRequest{
		Amount:   decimal.NewFromInt(100),
		Category: "food",
		Type:     "expense",
	}
	got := req.ToModel()
	if got.Currency != "RSD" {
		t.Errorf("expected currency RSD, got %q", got.Currency)
	}
}
