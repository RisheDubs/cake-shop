package store

import "cakeshop/api/internal/cakes"

func qar(amount int64) cakes.Money { return cakes.Money{Amount: amount * 100, Currency: "QAR"} }

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
				{ID: "6in", Label: "6 inch", Serves: 8, Price: qar(150)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: qar(220)},
				{ID: "10in", Label: "10 inch", Serves: 22, Price: qar(310)},
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
				{ID: "6in", Label: "6 inch", Serves: 8, Price: qar(160)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: qar(235)},
			},
			ImageURL:      "/images/red-velvet.jpg",
			LeadTimeHours: 24,
			Available:     true,
		},
		{
			ID:          "pistachio-rose",
			Name:        "Pistachio Rose",
			Description: "Pistachio sponge, rose-water cream and crushed pistachios.",
			Occasions:   []string{"eid", "wedding", "birthday"},
			Tags:        []string{"eggless"},
			Ingredients: []string{"flour", "pistachios", "rose water", "condensed milk", "cream", "sugar"},
			Allergens:   []string{"gluten", "dairy", "nuts"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: qar(180)},
				{ID: "8in", Label: "8 inch", Serves: 14, Price: qar(260)},
				{ID: "2tier", Label: "2-tier", Serves: 40, Price: qar(650)},
			},
			ImageURL:      "/images/pistachio-rose.jpg",
			LeadTimeHours: 48,
			Available:     true,
		},
		{
			ID:          "mango-seasonal",
			Name:        "Mango Cream (seasonal)",
			Description: "Fresh Alphonso mango and whipped cream. Summer only.",
			Occasions:   []string{"birthday"},
			Tags:        []string{"seasonal"},
			Ingredients: []string{"flour", "mango", "eggs", "cream", "sugar"},
			Allergens:   []string{"gluten", "eggs", "dairy"},
			Sizes: []cakes.Size{
				{ID: "6in", Label: "6 inch", Serves: 8, Price: qar(170)},
			},
			ImageURL:      "/images/mango.jpg",
			LeadTimeHours: 24,
			Available:     false,
		},
	}
}
