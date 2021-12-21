package store

import "cost-calculator/models"

type Store interface {
	InsertIncoming(incoming *models.Incoming) error
	GetIncoming(id int64) (*models.Incoming, error)
	UpdateIncoming(incoming *models.Incoming) error
	DeleteIncoming(id int64) error
	InsertSpends(spends *models.Spend) error
	GetSpend(id int64) (*models.Spend, error)
	UpdateSpend(spends *models.Spend) error
	DeleteSpend(id int64) error
}
