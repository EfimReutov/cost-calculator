package postgres

import (
	"cost-calculator/models"
	"cost-calculator/store"
)

func (p *Postgres) InsertCategory(category *models.Category) error {
	if category == nil {
		return store.ErrNilModel
	}
	if err := category.Validation(); err != nil {
		return err
	}
	sqlStatement := `
INSERT INTO categories (name)
VALUES ($1) RETURNING id`
	rows, err := p.db.Query(sqlStatement, category.Name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&category.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetCategory(id int64) (*models.Category, error) {
	sqlStatement := `
SELECT name FROM categories WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Category)
	row.ID = id
	if !rows.Next() {
		return nil, store.ErrNoRow
	} else {
		if err := rows.Scan(&row.Name); err != nil {
			return nil, err
		}
		return row, nil
	}
}

func (p *Postgres) UpdateCategory(category *models.Category) error {
	if category == nil {
		return store.ErrNilModel
	}
	if err := category.Validation(); err != nil {
		return err
	}
	if category.ID == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
UPDATE categories
SET name = $2
WHERE id = $1;`
	err := p.exec(sqlStatement, category.ID, category.Name)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteCategory(id int64) error {
	if id == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
DELETE FROM categories
WHERE id = $1`
	err := p.exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
