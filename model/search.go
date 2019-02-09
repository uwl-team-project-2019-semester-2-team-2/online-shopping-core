package model

type Search struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Variant string `json:"variant" db:"colour"`
}
