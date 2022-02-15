package postgres

import (
	"cost-calculator/models"
	"cost-calculator/store"
)

func (p *Postgres) InsertIncome(income *models.Income) error {
	if income == nil {
		return store.ErrNilModel
	}
	if err := income.Validation(); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO incoming (source_id, amount, date) VALUES ($1, $2, $3) RETURNING id`
	rows, err := p.db.Query(sqlStatement, income.SourceID, income.Amount, income.Date)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&income.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetIncome(id int64) (*models.Income, error) {
	sqlStatement := `SELECT source_id, amount, date FROM incoming WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Income)
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

func (p *Postgres) GetIncoming(page, limit int) ([]models.Income, error) {
	sqlStatement := `SELECT * FROM incoming OFFSET $1 LIMIT $2`
	rows, err := p.db.Query(sqlStatement, limit*(page-1), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]models.Income, limit)
	i := 0
	for rows.Next() {
		var row models.Income
		if err := rows.Scan(&row.ID, &row.SourceID, &row.Amount, &row.Date); err != nil {
			return nil, err
		}
		res[i] = row
		i++
	}
	return res[:i], nil
}

func (p *Postgres) UpdateIncome(income *models.Income) error {
	if income == nil {
		return store.ErrNilModel
	}
	if err := income.Validation(); err != nil {
		return err
	}
	if income.ID == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `UPDATE incoming SET source_id = $2, amount = $3 WHERE id = $1;`
	err := p.exec(sqlStatement, income.ID, income.SourceID, income.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteIncome(id int64) error {
	if id == 0 {
		return store.ErrZeroId
	}
	sqlStatement := `DELETE FROM incoming WHERE id = $1`
	err := p.exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
