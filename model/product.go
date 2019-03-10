package model

type Brand struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type Picture struct {
	URL           string     `json:"url" db:"url"`
}

type Product struct {
	Id            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Brand         string    `json:"brand" db:"brand_name"`
	Description   string    `json:"description" db:"description"`
	Stock         []Stock   `json:"stock"`
	Pictures      []Picture `json:"pictures"`
	Related       []Related `json:"variants"`
}

type Related struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"colour"`
}

type Stock struct {
	Size     string `json:"size" db:"size"`
	Quantity int    `json:"stock" db:"quantity"`
}
