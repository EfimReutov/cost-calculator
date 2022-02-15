package postgres

import (
	"cost-calculator/models"
	"cost-calculator/store"
)

func (p *Postgres) InsertSpend(spend *models.Spend) error {
	if spend == nil {
		return store.ErrNilModel
	}
	if err := spend.Validation(); err != nil {
		return err
	}
	sqlStatement := `
INSERT INTO spends (category_id, amount, description, date)
VALUES ($1, $2, $3, $4) RETURNING id`
	rows, err := p.db.Query(sqlStatement, spend.CategoryID, spend.Amount, spend.Description, spend.Date)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&spend.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetSpend(id int64) (*models.Spend, error) {
	sqlStatement := `
SELECT category_id, amount,description, date FROM spends WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Spend)
	row.ID = id
	if !rows.Next() {
		return nil, store.ErrNoRow
	} else {
		if err := rows.Scan(&row.CategoryID, &row.Amount, &row.Description, &row.Date); err != nil {
			return nil, err
		}
		return row, nil
	}
}

func (p *Postgres) GetSpends(page, limit int) ([]models.Spend, error) {
	sqlStatement := `SELECT * FROM spends OFFSET $1 LIMIT $2`
	rows, err := p.db.Query(sqlStatement, limit*(page-1), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]models.Spend, limit)
	i := 0
	for rows.Next() {
		var row models.Spend
		if err := rows.Scan(&row.ID, &row.CategoryID, &row.Amount, &row.Description, &row.Date); err != nil {
			return nil, err
		}
		res[i] = row
		i++
	}
	return res[:i], nil
}

func (p *Postgres) UpdateSpend(spend *models.Spend) error {
	if spend == nil {
		return store.ErrNilModel
	}
	if err := spend.Validation(); err != nil {
		return err
	}
	if spend.ID == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
UPDATE spends
SET category_id = $2, amount = $3, description = $4  
WHERE id = $1;`
	err := p.exec(sqlStatement, spend.ID, spend.CategoryID, spend.Amount, spend.Description)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteSpend(id int64) error {
	if id == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `
DELETE FROM spends
WHERE id = $1`
	err := p.exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
