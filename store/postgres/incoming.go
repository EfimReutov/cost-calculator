package postgres

import (
	"cost-calculator/models"
	"cost-calculator/store"
)

func (p *Postgres) InsertIncoming(incoming *models.Incoming) error {
	if incoming == nil {
		return store.ErrNilModel
	}
	if err := incoming.Validation(); err != nil {
		return err
	}

	sqlStatement := `
INSERT INTO incoming (source_id, amount, date)
VALUES ($1, $2, $3) RETURNING id`
	rows, err := p.db.Query(sqlStatement, incoming.SourceID, incoming.Amount, incoming.Date)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&incoming.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetIncoming(id int64) (*models.Incoming, error) {
	sqlStatement := `
SELECT source_id, amount, date FROM incoming WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Incoming)
	row.ID = id
	if !rows.Next() {
		return nil, store.ErrNoRow
	} else {
		if err := rows.Scan(&row.SourceID, &row.Amount, &row.Date); err != nil {
			return nil, err
		}
		return row, nil
	}
}

func (p *Postgres) UpdateIncoming(incoming *models.Incoming) error {
	if incoming == nil {
		return store.ErrNilModel
	}
	if err := incoming.Validation(); err != nil {
		return err
	}
	if incoming.ID == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
UPDATE incoming
SET source_id = $2, amount = $3 
WHERE id = $1;`
	err := p.exec(sqlStatement, incoming.ID, incoming.SourceID, incoming.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteIncoming(id int64) error {
	if id == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
DELETE FROM incoming
WHERE id = $1`
	err := p.exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
