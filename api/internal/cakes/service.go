package cakes

import (
	"fmt"
	"slices"
)

// Service contains catalogue business logic. Handlers and AI tools call this,
// never the repository directly.
type Service struct {
	repo Repository
}

// NewService builds a Service on top of any Repository.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// List returns cakes matching the filter. Invalid filters return an error wrapping ErrInvalid.
func (s *Service) List(f Filter) ([]Cake, error) {
	f = f.Normalize()
	if err := f.Validate(); err != nil {
		return nil, err
	}
	all, err := s.repo.ListCakes()
	if err != nil {
		return nil, err
	}
	out := make([]Cake, 0, len(all))
	for _, c := range all {
		if f.OnlyAvailable && !c.Available {
			continue
		}
		if f.Occasion != "" && !slices.Contains(c.Occasions, f.Occasion) {
			continue
		}
		if f.Tag != "" && !slices.Contains(c.Tags, f.Tag) {
			continue
		}
		out = append(out, c)
	}
	return out, nil
}

// Get returns one cake by id.
func (s *Service) Get(id string) (Cake, error) {
	return s.repo.GetCake(id)
}

// Quote is the single source of truth for what an order line costs.
// The AI assistant's "quote" tool calls exactly this function.
type Quote struct {
	CakeID   string `json:"cakeId"`
	CakeName string `json:"cakeName"`
	Size     Size   `json:"size"`
	Quantity int    `json:"quantity"`
	Total    Money  `json:"total"`
}

// Quote prices a cake/size/quantity combination.
func (s *Service) Quote(cakeID, sizeID string, qty int) (Quote, error) {
	if qty < 1 || qty > 20 {
		return Quote{}, fmt.Errorf("quantity must be between 1 and 20")
	}
	c, err := s.repo.GetCake(cakeID)
	if err != nil {
		return Quote{}, err
	}
	if !c.Available {
		return Quote{}, ErrUnavailable
	}
	size, err := c.SizeByID(sizeID)
	if err != nil {
		return Quote{}, err
	}
	return Quote{
		CakeID:   c.ID,
		CakeName: c.Name,
		Size:     size,
		Quantity: qty,
		Total:    Money{Amount: size.Price.Amount * int64(qty), Currency: size.Price.Currency},
	}, nil
}