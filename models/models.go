package models

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
	"unicode"
)

// Category represents model for 'categories' table.
type Category struct {
	ID   int64
	Name string
}

// Source represents model for 'sources' table.
type Source struct {
	ID   int64
	Name string
	Date time.Time
}

// Income represents model for 'incoming' table.
type Income struct {
	ID       int64           `json:"id"`
	SourceID int32           `json:"source_id,omitempty"`
	Amount   decimal.Decimal `json:"amount"`
	Date     time.Time       `json:"date"`
}

// Spend represents model for 'spends' table.
type Spend struct {
	ID          int64           `json:"id"`
	CategoryID  int32           `json:"category_id"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (c *Category) Validation() error {
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}
	for _, v := range []rune(c.Name) {
		if unicode.IsDigit(v) {
			return errors.New("name cannot contain numbers")
		}
	}
	return nil
}

func (s *Source) Validation() error {
	if s.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (i *Income) Validation() error {
	if i.SourceID == 0 {
		return errors.New("source id cannot be zero")
	}
	if i.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount cannot be equal or less than zero")
	}
	return nil
}

func (s *Spend) Validation() error {
	if s.CategoryID == 0 {
		return errors.New("category id cannot be zero")
	}
	if s.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount cannot be equal or less than zero")
	}
	return nil
}
