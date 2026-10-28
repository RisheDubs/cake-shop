package store

import "cakeshop/api/internal/cakes"

// inr converts whole rupees to paise.
func inr(rupees int64) cakes.Money { return cakes.Money{Amount: rupees * 100, Currency: "INR"} }

// SeedCakes is sample data for local development and tests.
// Replace with the client's real menu (and photos) during Sprint 2.
func SeedCakes() []cakes.Cake {
	return []cakes.Cake{
		{
			ID:          "chocolate-truffle",
			Name:        "Chocolate Truffle",
			Description: "Three layers of dark chocolate sponge with silky ganache.",
			Occasions:   []string{"birthday", "anniversary"},
			Tags:        []string{"bestseller"},
			Ingredients: []string{"dark chocolate", "cocoa", "flour", "eggs", "butter", "cream", "sugar"},
			Allergens:   []string{"gluten", "eggs", "dairy"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: inr(850)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: inr(1250)},
				{ID: "10in", Label: "10 inch", Serves: 22, Price: inr(1750)},
			},
			ImageURL:      "/images/chocolate-truffle.jpg",
			LeadTimeHours: 24,
			Available:     true,
		},
		{
			ID:          "red-velvet",
			Name:        "Red Velvet",
			Description: "Soft red velvet layers with cream-cheese frosting.",
			Occasions:   []string{"birthday", "anniversary", "valentine"},
			Ingredients: []string{"flour", "buttermilk", "cocoa", "eggs", "cream cheese", "butter", "sugar"},
			Allergens:   []string{"gluten", "eggs", "dairy"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: inr(950)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: inr(1400)},
			},
			ImageURL:      "/images/red-velvet.jpg",
			LeadTimeHours: 24,
			Available:     true,
		},
		{
			ID:          "pistachio-rose",
			Name:        "Pistachio Rose",
			Description: "Pistachio sponge, rose-water cream and crushed pistachios.",
			Occasions:   []string{"diwali", "wedding", "birthday"},
			Tags:        []string{"eggless"},
			Ingredients: []string{"flour", "pistachios", "rose water", "condensed milk", "cream", "sugar"},
			Allergens:   []string{"gluten", "dairy", "nuts"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: inr(1100)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: inr(1600)},
				{ID: "2tier", Label: "2-tier", Serves: 40, Price: inr(3800)},
			},
			ImageURL:      "/images/pistachio-rose.jpg",
			LeadTimeHours: 48,
			Available:     true,
		},
		{
			ID:          "mango-seasonal",
			Name:        "Mango Cream (seasonal)",
			Description: "Fresh Alphonso mango and whipped cream. April to June only.",
			Occasions:   []string{"birthday"},
			Tags:        []string{"seasonal"},
			Ingredients: []string{"flour", "mango", "eggs", "cream", "sugar"},
			Allergens:   []string{"gluten", "eggs", "dairy"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: inr(900)},
			},
			ImageURL:      "/images/mango.jpg",
			LeadTimeHours: 24,
			Available:     false,
		},
	}
}
