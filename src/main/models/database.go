package models

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

func (d *Database) Init() {
	db, err := sql.Open("mysql", "jj:ranger@localhost:3306/shop")

	if err != nil {
		fmt.Print(err)
	}

	d.DB = db
}
