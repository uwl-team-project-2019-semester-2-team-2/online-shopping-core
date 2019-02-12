package database

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

var (
	DB *DBType
)

type DBType struct {
	db *sql.DB
}

func (db *DBType) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func OpenMysql(user, password, addr, dbName string) (*DBType, error) {
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

	DB = &DBType{
		db: db,
	}

	return DB, nil
}
