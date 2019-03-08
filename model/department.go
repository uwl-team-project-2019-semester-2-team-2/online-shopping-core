package model

type Department struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
}
