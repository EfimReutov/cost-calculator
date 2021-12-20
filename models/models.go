package models

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
	"unicode"
)

type Category struct {
	ID   int8
	Name string
}

type Source struct {
	ID   int8
	Name string
	Date time.Time
}

type Incoming struct {
	ID       int64           `json:"id"`
	SourceID int8            `json:"source_id,omitempty"`
	Amount   decimal.Decimal `json:"amount"`
	Date     time.Time       `json:"date"`
}

type Spend struct {
	ID          int64           `json:"id"`
	CategoryID  int8            `json:"category_id"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description"`
	Date        time.Time       `json:"date"`
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

func (i *Incoming) Validation() error {
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