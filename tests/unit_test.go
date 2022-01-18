package tests

import (
	"cost-calculator/models"
	"errors"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestCategory_Validation(t *testing.T) {
	tests := []struct {
		name     string
		category models.Category
		err      error
	}{
		{
			name: "all good",
			category: models.Category{
				Name: "base_service",
			},
		},
		{
			name: "name is empty",
			category: models.Category{
				Name: "",
			},
			err: errors.New("name cannot be empty"),
		},
		{
			name: "name contain numbers",
			category: models.Category{
				Name: "2222",
			},
			err: errors.New("name cannot contain numbers"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.category.Validation()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Result: %v not equal expect value: %v", err, tt.err)
			}
		})
	}
}

func TestSources_Validation(t *testing.T) {
	tests := []struct {
		name   string
		source models.Source
		err    error
	}{
		{
			name: "all good",
			source: models.Source{
				Name: "base_service",
			},
		},
		{
			name: "name is empty",
			source: models.Source{
				Name: "",
			},
			err: errors.New("name cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.source.Validation()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Result: %v not equal expect value: %v", err, tt.err)
			}
		})
	}
}

func TestIncoming_Validation(t *testing.T) {
	tests := []struct {
		name     string
		incoming models.Incoming
		err      error
	}{
		{
			name: "all good",
			incoming: models.Incoming{
				SourceID: 1,
				Amount:   decimal.New(2, 2),
			},
		},
		{
			name: "source id is zero",
			incoming: models.Incoming{
				Amount: decimal.New(2, 0),
			},
			err: errors.New("source id cannot be zero"),
		},
		{
			name: "amount is zero",
			incoming: models.Incoming{
				SourceID: 2,
			},
			err: errors.New("amount cannot be equal or less than zero"),
		},
		{
			name: "amount is negative",
			incoming: models.Incoming{
				SourceID: 2,
				Amount:   decimal.New(-2, -2),
			},
			err: errors.New("amount cannot be equal or less than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.incoming.Validation()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Result: %v not equal expect value: %v", err, tt.err)
			}
		})
	}
}

func TestSpend_Validation(t *testing.T) {
	tests := []struct {
		name  string
		spend models.Spend
		err   error
	}{
		{
			name: "all good",
			spend: models.Spend{
				CategoryID: 1,
				Amount:     decimal.New(2, 2),
			},
		},
		{
			name: "category id is zero",
			spend: models.Spend{
				Amount: decimal.New(2, 0),
			},
			err: errors.New("category id cannot be zero"),
		},
		{
			name: "amount is zero",
			spend: models.Spend{
				CategoryID: 2,
			},
			err: errors.New("amount cannot be equal or less than zero"),
		},
		{
			name: "amount is negative",
			spend: models.Spend{
				CategoryID: 2,
				Amount:     decimal.New(-2, -2),
			},
			err: errors.New("amount cannot be equal or less than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.spend.Validation()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Result: %v not equal expect value: %v", err, tt.err)
			}
		})
	}
}
