package cakes

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ErrInvalid marks input that breaks a rule. Handlers turn it into HTTP 400.
var ErrInvalid = errors.New("invalid input")

// slug is the allowed shape for occasions and tags: lowercase letters, digits and
// hyphens, starting with a letter, at most 30 characters. e.g. "birthday", "sugar-free".
var slug = regexp.MustCompile(`^[a-z][a-z0-9-]{0,29}$`)

// Normalize trims spaces and lowercases the filter, so "Diwali " matches "diwali".
func (f Filter) Normalize() Filter {
	f.Occasion = strings.ToLower(strings.TrimSpace(f.Occasion))
	f.Tag = strings.ToLower(strings.TrimSpace(f.Tag))
	return f
}

// Validate reports the first rule the filter breaks. Empty fields are allowed (no filter).
func (f Filter) Validate() error {
	if f.Occasion != "" && !slug.MatchString(f.Occasion) {
		return fmt.Errorf("%w: occasion %q must be lowercase letters, digits or '-', max 30", ErrInvalid, f.Occasion)
	}
	if f.Tag != "" && !slug.MatchString(f.Tag) {
		return fmt.Errorf("%w: tag %q must be lowercase letters, digits or '-', max 30", ErrInvalid, f.Tag)
	}
	return nil
}
