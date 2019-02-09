package database

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

type DB struct {
	db *sql.DB
}

func OpenMysql(user, password, addr, dbName string) (*DB, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = user
	mysqlConfig.Passwd = password
	mysqlConfig.Addr = addr
	mysqlConfig.DBName = dbName

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		return nil, errors.New("unable to open database connection:" + err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("unable to open database connection:" + err.Error())
	}

	return &DB{
		db: db,
	}, nil
}
