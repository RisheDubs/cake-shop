package cakes_test

import (
	"errors"
	"testing"

	"cakeshop/api/internal/cakes"
	"cakeshop/api/internal/store"
)

func newService() *cakes.Service {
	return cakes.NewService(store.NewMemory(store.SeedCakes()))
}

func TestListFiltersByTag(t *testing.T) {
	got, err := newService().List(cakes.Filter{Tag: "eggless"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatal("expected at least one eggless cake in seed data")
	}
	for _, c := range got {
		found := false
		for _, tag := range c.Tags {
			if tag == "eggless" {
				found = true
			}
		}
		if !found {
			t.Errorf("%s is not eggless", c.ID)
		}
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		name    string
		cake    string
		size    string
		qty     int
		want    int64
		wantErr error
	}{
		{"single 6in chocolate", "chocolate-truffle", "6in", 1, 85000, nil},
		{"two 8in chocolate", "chocolate-truffle", "8in", 2, 250000, nil},
		{"unknown cake", "nope", "6in", 1, 0, cakes.ErrNotFound},
		{"unknown size", "chocolate-truffle", "99in", 1, 0, cakes.ErrNotFound},
		{"unavailable cake", "mango-seasonal", "6in", 1, 0, cakes.ErrUnavailable},
	}
	svc := newService()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q, err := svc.Quote(tc.cake, tc.size, tc.qty)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("want error %v, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if q.Total.Amount != tc.want {
				t.Errorf("total = %d, want %d", q.Total.Amount, tc.want)
			}
		})
	}
}

func TestQuoteRejectsBadQuantity(t *testing.T) {
	if _, err := newService().Quote("chocolate-truffle", "6in", 0); err == nil {
		t.Fatal("expected error for quantity 0")
	}
}

func TestListFiltersByOccasion(t *testing.T) {
	got, err := newService().List(cakes.Filter{Occasion: "diwali"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].ID != "pistachio-rose" {
		t.Fatalf("want only pistachio-rose, got %d cakes: %+v", len(got), got)
	}
}

func TestListOnlyAvailable(t *testing.T) {
	svc := newService()

	all, _ := svc.List(cakes.Filter{})
	onSale, _ := svc.List(cakes.Filter{OnlyAvailable: true})

	if len(all) != 4 {
		t.Errorf("all cakes = %d, want 4", len(all))
	}
	if len(onSale) != 3 {
		t.Errorf("cakes on sale = %d, want 3 (mango is unavailable)", len(onSale))
	}
	for _, c := range onSale {
		if c.ID == "mango-seasonal" {
			t.Error("mango-seasonal should be hidden when OnlyAvailable is true")
		}
	}
}
