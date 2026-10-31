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
	if q.Total.Amount != 280000 { // 2 x ₹1,400
		t.Errorf("total = %d, want 280000", q.Total.Amount)
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

func TestGetCakeReturnsSizes(t *testing.T) {
	rec := do(t, newServer(), "GET", "/v1/cakes/pistachio-rose")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var c cakes.Cake
	if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
		t.Fatal(err)
	}
	if len(c.Sizes) != 3 {
		t.Errorf("sizes = %d, want 3", len(c.Sizes))
	}
}

func TestQuoteEndpointRejectsBadQuantity(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"not a number", "/v1/cakes/red-velvet/quote?size=8in&qty=abc"},
		{"zero", "/v1/cakes/red-velvet/quote?size=8in&qty=0"},
		{"too many", "/v1/cakes/red-velvet/quote?size=8in&qty=50"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if rec := do(t, newServer(), "GET", tc.path); rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want 400", rec.Code)
			}
		})
	}
}
