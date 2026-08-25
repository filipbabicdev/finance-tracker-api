package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootReturnsRegisteredRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewTransactionHandler(nil)
	SetupRoutes(r, h)
	r.GET("/", RootHandler(r, "test"))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !strings.Contains(w.Body.String(), "/transactions") {
		t.Errorf("expected response to list /transactions route, got: %s", w.Body.String())
	}
}

func TestRootIncludesRoutesRegisteredAfterRoot(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/", RootHandler(r, "test")) // root registered FIRST

	h := NewTransactionHandler(nil)
	SetupRoutes(r, h) // transaction routes registered AFTER root

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !strings.Contains(w.Body.String(), "/transactions") {
		t.Errorf("expected response to include routes registered after root, got: %s", w.Body.String())
	}
}

func TestRootReturns200AndJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/", RootHandler(r, "test"))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		t.Errorf("expected application/json content type, got %q", ct)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Errorf("expected valid JSON body, got error: %v", err)
	}
}

func TestNoRouteReturns404JSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.NoRoute(NoRouteHandler())

	req := httptest.NewRequest(http.MethodGet, "/does-not-exist", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		t.Errorf("expected application/json content type, got %q", ct)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON body, got error: %v", err)
	}
	if _, ok := body["error"]; !ok {
		t.Errorf("expected body to have an \"error\" key, got: %s", w.Body.String())
	}
}
