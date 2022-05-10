package store

import "cost-calculator/models"

// StorePostgres describes all the methods required by the server to work with the DB.
type StorePostgres interface {
	InsertCategory(category *models.Category) error
	GetCategory(id int64) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int64) error
	InsertSource(source *models.Source) error
	GetSource(id int64) (*models.Source, error)
	UpdateSource(sSource *models.Source) error
	DeleteSource(id int64) error
	InsertIncome(income *models.Income) error
	GetIncome(id int64) (*models.Income, error)
	GetIncoming(page, limit int) ([]models.Income, error)
	UpdateIncome(income *models.Income) error
	DeleteIncome(id int64) error
	InsertSpend(spend *models.Spend) error
	GetSpend(id int64) (*models.Spend, error)
	GetSpends(page, limit int) ([]models.Spend, error)
	UpdateSpend(spend *models.Spend) error
	DeleteSpend(id int64) error
}

type StoreRedis interface {
	InsertOTP()
	GetOTP()
	DeleteOTP()
}
