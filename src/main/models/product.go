package models

type Brand struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	Description string	`json:"description"`	
}

type Product struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	Brand string		`json:"brand"`
	Colour string		`json:"colour"`
	Description string	`json:"description"`
	Stock []Stock		`json:"stock"`
	Pictures []string	`json:"pictures"`
	Related []Related	`json:"related"`
}

type Related struct {
	Id int				`json:"id"`
	Title string		`json:"title"`
}

type Size struct {
	Size string			`json:"size"`
	Quantity int		`json:"stock"`
}

type Stock struct {
	Size string			`json:"size"`
	Quantity int		`json:"stock"`
}