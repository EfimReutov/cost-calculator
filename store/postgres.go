package store

import (
	"cost-calculator/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	db *sql.DB
}

func NewDB(cfg *Config) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.connStr())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected successful")
	return &Postgres{db: db}, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}

func (p *Postgres) exec(query string, args ...interface{}) error {
	result, err := p.db.Exec(query, args...)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		return ErrNoRow
	}
	return nil
}

func (p *Postgres) InsertIncoming(incoming *models.Incoming) error {
	if incoming == nil {
		return ErrNilModel
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
		return nil, ErrNoRow
	} else {
		if err := rows.Scan(&row.SourceID, &row.Amount, &row.Date); err != nil {
			return nil, err
		}
		return row, nil
	}
}

func (p *Postgres) UpdateIncoming(incoming *models.Incoming) error {
	if incoming == nil {
		return ErrNilModel
	}
	if err := incoming.Validation(); err != nil {
		return err
	}
	if incoming.ID == 0 {
		return ErrZeroId
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
		return ErrZeroId
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

func (p *Postgres) InsertSpends(spends *models.Spend) error {
	if spends == nil {
		return ErrNilModel
	}
	if err := spends.Validation(); err != nil {
		return err
	}
	sqlStatement := `
INSERT INTO spends (category_id, amount, description, date)
VALUES ($1, $2, $3, $4) RETURNING id`
	rows, err := p.db.Query(sqlStatement, spends.CategoryID, spends.Amount, spends.Description, spends.Date)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&spends.ID); err != nil {
			return err
		}
	}
	return nil
}

func (p *Postgres) GetSpend(id int64) (*models.Spend, error) {
	sqlStatement := `
SELECT category_id, amount, description, date FROM spends WHERE id = $1`
	rows, err := p.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	row := new(models.Spend)
	row.ID = id
	if !rows.Next() {
		return nil, ErrNoRow
	} else {
		if err := rows.Scan(&row.CategoryID, &row.Amount, &row.Description, &row.Date); err != nil {
			return nil, err
		}
		return row, nil
	}
}

func (p *Postgres) UpdateSpend(spends *models.Spend) error {
	if spends == nil {
		return ErrNilModel
	}
	if err := spends.Validation(); err != nil {
		return err
	}
	if spends.ID == 0 {
		return ErrZeroId
	}
	sqlStatement := `
UPDATE spends
SET category_id = $2, amount = $3, description = $4  
WHERE id = $1;`
	err := p.exec(sqlStatement, spends.ID, spends.CategoryID, spends.Amount, spends.Description)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteSpend(id int64) error {
	if id == 0 {
		return ErrZeroId
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
