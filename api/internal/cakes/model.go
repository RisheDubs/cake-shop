// Package cakes holds the cake catalogue: the domain model, business rules
// (pricing, availability) and the repository interface the storage layer implements.
//
// Rule of the project: prices and ingredients live ONLY here and in the database.
// The website, the mobile app and the AI assistant all read them through this package,
// so nobody (including the AI) can invent a price.
package cakes

import "errors"

// ErrNotFound is returned when a cake or size does not exist.
var ErrNotFound = errors.New("not found")

// ErrUnavailable is returned when a cake exists but is not currently on sale.
var ErrUnavailable = errors.New("cake not available")

// Money is stored in minor units (dirhams for QAR: 1 QAR = 100 dirhams)
// so we never have floating-point rounding problems.
type Money struct {
	Amount   int64  `json:"amount"`   // e.g. 18500 = 185.00 QAR
	Currency string `json:"currency"` // ISO 4217, e.g. "QAR"
}

// Size is one orderable size of a cake, with its own price.
type Size struct {
	ID     string `json:"id"`     // stable id, e.g. "6in"
	Label  string `json:"label"`  // shown to customers, e.g. "6 inch"
	Serves int    `json:"serves"` // approximate number of people
	Price  Money  `json:"price"`
}

// Cake is a product on the menu.
type Cake struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Occasions     []string `json:"occasions"`   // birthday, wedding, eid ...
	Tags          []string `json:"tags"`        // eggless, sugar-free ...
	Ingredients   []string `json:"ingredients"` // shown on hover / tap
	Allergens     []string `json:"allergens"`   // gluten, dairy, eggs, nuts ...
	Sizes         []Size   `json:"sizes"`
	ImageURL      string   `json:"imageUrl"`
	LeadTimeHours int      `json:"leadTimeHours"` // minimum notice before delivery
	Available     bool     `json:"available"`
}

// SizeByID returns the size with the given id.
func (c Cake) SizeByID(id string) (Size, error) {
	for _, s := range c.Sizes {
		if s.ID == id {
			return s, nil
		}
	}
	return Size{}, ErrNotFound
}

// Filter narrows a catalogue listing. Empty fields mean "no filter".
type Filter struct {
	Occasion      string
	Tag           string
	OnlyAvailable bool
}

// Repository is implemented by the storage layer (in-memory now, DynamoDB in Sprint 2).
type Repository interface {
	ListCakes() ([]Cake, error)
	GetCake(id string) (Cake, error)
}
