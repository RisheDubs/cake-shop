package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cakeshop/api/internal/cakes"
	"cakeshop/api/internal/httpapi"
	"cakeshop/api/internal/store"
)

func newServer() http.Handler {
	svc := cakes.NewService(store.NewMemory(store.SeedCakes()))
	return httpapi.NewRouter(httpapi.Deps{Cakes: svc, AllowedOrigins: []string{"http://localhost:3000"}})
}

func do(t *testing.T, h http.Handler, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
	return rec
}

func TestHealth(t *testing.T) {
	if rec := do(t, newServer(), "GET", "/health"); rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestListHidesUnavailable(t *testing.T) {
	rec := do(t, newServer(), "GET", "/v1/cakes")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var body struct{ Cakes []cakes.Cake }
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	for _, c := range body.Cakes {
		if !c.Available {
			t.Errorf("unavailable cake %q returned", c.ID)
		}
	}
}

func TestGetUnknownCake(t *testing.T) {
	if rec := do(t, newServer(), "GET", "/v1/cakes/does-not-exist"); rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}

func TestQuoteEndpoint(t *testing.T) {
	rec := do(t, newServer(), "GET", "/v1/cakes/red-velvet/quote?size=8in&qty=2")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", rec.Code, rec.Body)
	}
	var q cakes.Quote
	if err := json.NewDecoder(rec.Body).Decode(&q); err != nil {
		t.Fatal(err)
	}
	if q.Total.Amount != 47000 {
		t.Errorf("total = %d, want 47000", q.Total.Amount)
	}
}

func TestCORSPreflight(t *testing.T) {
	req := httptest.NewRequest("OPTIONS", "/v1/cakes", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	rec := httptest.NewRecorder()
	newServer().ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
		t.Errorf("allow-origin = %q", got)
	}
}
