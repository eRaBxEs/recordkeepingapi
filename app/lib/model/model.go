package model

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/shopspring/decimal"
)

// DatabaseInfo ...
type DatabaseInfo struct {
	User     string `sane:"user"`
	DBName   string `sane:"dbname"`
	Password string `sane:"password"`
}

// ConfigFile ...
type ConfigFile struct {
	Host   string       `sane:"host"`
	Port   string       `sane:"port"`
	DBInfo DatabaseInfo `sane:"dbinfo"`
}

// Income ...
type Income struct {
	ID          int64           `json:"id"`
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
	Time        time.Time       `json:"time"`
}

// Expense ...
type Expense struct {
	ID          int64           `json:"id"`
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
	Time        time.Time       `json:"time"`
}

// Save ...
func (s *Expense) Save(db *pg.DB) error {

	if err := db.Insert(s); err != nil {
		return err
	}

	return nil
}

// Save ...
func (s *Income) Save(db *pg.DB) error {

	if err := db.Insert(s); err != nil {
		return err
	}

	return nil
}

// GetAll ...
func (s *Expense) GetAll(db *pg.DB) ([]Expense, error) {

	expenses := []Expense{}

	err := db.Model(&expenses).Select()

	if err != nil {
		return expenses, err
	}

	return expenses, nil
}

// GetAll ...
func (s *Income) GetAll(db *pg.DB) ([]Income, error) {

	income := []Income{}

	err := db.Model(&income).Select()

	if err != nil {
		return income, err
	}

	return income, nil
}
