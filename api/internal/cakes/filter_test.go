package cakes_test

import (
	"errors"
	"strings"
	"testing"

	"cakeshop/api/internal/cakes"
)

func TestFilterValidate(t *testing.T) {
	tests := []struct {
		name    string
		filter  cakes.Filter
		wantErr bool
	}{
		{"empty filter is fine", cakes.Filter{}, false},
		{"simple occasion", cakes.Filter{Occasion: "birthday"}, false},
		{"hyphenated tag", cakes.Filter{Tag: "sugar-free"}, false},
		{"uppercase is normalized first", cakes.Filter{Occasion: " Diwali "}, false},
		{"spaces inside", cakes.Filter{Occasion: "new year"}, true},
		{"symbols", cakes.Filter{Tag: "egg<less>"}, true},
		{"starts with digit", cakes.Filter{Tag: "1st"}, true},
		{"too long", cakes.Filter{Tag: strings.Repeat("a", 31)}, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.filter.Normalize().Validate()
			if tc.wantErr {
				if !errors.Is(err, cakes.ErrInvalid) {
					t.Fatalf("want ErrInvalid, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("want no error, got %v", err)
			}
		})
	}
}

func TestListIsCaseInsensitive(t *testing.T) {
	got, err := newService().List(cakes.Filter{Occasion: "DIWALI"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 {
		t.Fatalf("want 1 cake, got %d", len(got))
	}
}