package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Db interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}

type Postgres struct {
	db *sql.DB
}

func CreateDb() (*Postgres, error) {

	connstr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connstr)
 
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (pg *Postgres) CreateAccount(*Account) error {
	return nil
}

func (pg *Postgres) DeleteAccount(id int) error {
	return nil
}

func (pg *Postgres) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
