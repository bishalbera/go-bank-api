package main

import "database/sql"

type Db interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}

type postgres struct {
	db *sql.DB
}


func createDb() (*postgres, error) {

	connstr:= "user=postgres dbname=postgres password=gobank ssl-mode=disable"
	db,err:= sql.Open("postgres",connstr)

	if err != nil {
		return nil, err
	}
	if err:= db.Ping(); err!= nil{
		return nil, err
	}

	return &postgres{
		db: db,
	}, nil
}