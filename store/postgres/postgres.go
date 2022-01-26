package postgres

import (
	"cost-calculator/config"
	"cost-calculator/store"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	db *sql.DB
}

// NewDB creates a connection to the postgres db.
func NewDB(cfg *config.Configuration) (*Postgres, error) {
	db, err := sql.Open("postgres", newConfig(cfg).connStr())
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

// Close closes the database and prevents new queries from starting.
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
		return store.ErrNoRow
	}
	return nil
}
