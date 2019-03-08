package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Connection *sqlx.DB
}

func Connect(conf *mysql.Config) (Database, error) {
	db, err := sqlx.Open("mysql", conf.FormatDSN())

	if err != nil {
		return Database{}, err
	}

	err = db.Ping()

	if err != nil {
		return Database{}, err
	}

	return Database{db}, nil
}

func (db *Database) Get(id string, dest interface{}, query string) error {
	if err := db.Connection.Select(dest, query, id); err != nil {
		return err
	}
	return nil
}

func (db *Database) Count(id string, dest *int, query string) error {
	if err := db.Connection.Get(dest, query, id); err != nil {
		return err
	}
	return nil
}

func (db *Database) List(dest interface{}, query string) error {
	if err := db.Connection.Select(dest, query); err != nil {
		return err
	}
	return nil
}

func (db *Database) Paginate(dest interface{}, query string, page int, perPage int) error {
	if err := db.Connection.Select(dest, query); err != nil {
		return err
	}
	return nil
}