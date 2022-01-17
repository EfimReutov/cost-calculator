package postgres

import (
	"cost-calculator/models"
	"cost-calculator/store"
)

func (p *Postgres) InsertSource(source *models.Source) error {
	if source == nil {
		return store.ErrNilModel
	}
	if err := source.Validation(); err != nil {
		return err
	}
	sqlStatement := `
INSERT INTO sources (name)
VALUES ($1) RETURNING id`
	rows, err := p.db.Query(sqlStatement, source.ID, source.Name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&source.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetSource(id int64) (*models.Source, error) {
	sqlStatement := `
SELECT name FROM sources WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Source)
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

func (p *Postgres) UpdateSource(source *models.Source) error {
	if source == nil {
		return store.ErrNilModel
	}
	if err := source.Validation(); err != nil {
		return err
	}
	if source.ID == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
UPDATE sources
SET name = $2  
WHERE id = $1;`
	err := p.exec(sqlStatement, source.ID, source.Name)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteSource(id int64) error {
	if id == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
DELETE FROM sources
WHERE id = $1`
	err := p.exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
