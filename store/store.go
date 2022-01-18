package store

import "C"
import "cost-calculator/models"

type Store interface {
	InsertCategory(category *models.Category) error
	GetCategory(id int64) (*models.Category, error)
	UpdateCategory(incoming *models.Category) error
	DeleteCategory(id int64) error
	InsertSource(source *models.Source) error
	GetSource(id int64) (*models.Source, error)
	UpdateSource(sSource *models.Source) error
	DeleteSource(id int64) error
	InsertIncoming(incoming *models.Incoming) error
	GetIncoming(id int64) (*models.Incoming, error)
	UpdateIncoming(incoming *models.Incoming) error
	DeleteIncoming(id int64) error
	InsertSpend(spends *models.Spend) error
	GetSpend(id int64) (*models.Spend, error)
	UpdateSpend(spends *models.Spend) error
	DeleteSpend(id int64) error
}
