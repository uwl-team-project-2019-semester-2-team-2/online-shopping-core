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

func (db *Database) Query(query string, queryMap map[string]interface{}) (*sqlx.Rows, error) {

	query, args, err := sqlx.Named(query, queryMap)
	query, args, err = sqlx.In(query, args...)
	query = db.Connection.Rebind(query)
	rows, err := db.Connection.Queryx(query, args...)

	if err != nil {
		return &sqlx.Rows{}, err
	}
	return rows, nil
}

//	for rows.Next() {
//		switch reflect.Indirect(reflect.ValueOf(dest)).Kind() {
//			case reflect.Slice:
//				value := reflect.TypeOf(dest)
//				t := value.Elem()
//				vp := reflect.New(t)
//				v := reflect.Indirect(vp)
//
//
//				fmt.Println(reflect.ValueOf(v))
//
//				err := rows.StructScan(&v)
//
//				if err != nil {
//					return err
//				}
//
//				dest = append([]interface{}{dest}, v)
//			case reflect.Int:
//				err := rows.Scan(dest)
//
//				if err != nil {
//					return err
//				}
//		}
//	}
//	return nil
//}

func (db *Database) Scan(dest interface{}, rows *sqlx.Rows) error {
	rows.Next()
	err := rows.Scan(dest)

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Get(dest interface{}, query string, id string) error {
	if err := db.Connection.Select(dest, query, id); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetOne(id string, dest interface{}, query string) error {
	if err := db.Connection.Get(dest, query, id); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetSlice(dest interface{}, query string, args ...interface{}) error {
	if err := db.Connection.Select(dest, query, args...); err != nil {
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
