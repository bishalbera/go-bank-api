package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
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

func (pg *Postgres) Init() error {
	return pg.createAccountTable()
}

func (pg *Postgres) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		balance float,
		first_name varchar(50),
		last_name varchar(50),
		number int,
		created_at timestamp
	)`
	_, err := pg.db.Exec(query)
	return err
}

func (pg *Postgres) CreateAccount(acc *Account) error {
	sqlQuery := `INSERT INTO account (
		first_name, 
		last_name,
		balance,
		number, 
		created_at
	) VALUES($1, $2, $3, $4, $5)`

	_, err := pg.db.Query(sqlQuery, acc.FirstName, acc.LastName, acc.Balance, acc.Number, acc.CreatedAt)
	if err != nil {
		panic(err)
	}
	return nil
}

func (pg *Postgres) DeleteAccount(id int) error {
	query := "DELETE FROM account WHERE id = $1"

	_, err := pg.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("account not found with id %d", id)
	}

	return nil
}

func (pg *Postgres) GetAccountByID(id int) (*Account, error) {
	query := "SELECT id, first_name, last_name, balance, number, created_at FROM account WHERE id = $1"

	acc, err := getRows(pg, query, id)
	if err != nil {
		return nil, err
	}
	if len(acc) == 0 {
		return nil, fmt.Errorf("no account found with id %d", id)
	}

	return acc[0], nil
}

func (pg *Postgres) GetAccounts() ([]*Account, error) {
	query := "SELECT id, first_name, last_name, balance, number, created_at FROM account"

	return getRows(pg, query)

}

func getRows(pg *Postgres, query string, args ...interface{}) ([]*Account, error) {

	rows, err := pg.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Balance,
			&account.Number,
			&account.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, err
}
