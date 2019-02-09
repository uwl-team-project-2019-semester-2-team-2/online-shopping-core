package model

//type Product struct {
//	Id          int         `json:"id"`
//	Name        string      `json:"name"`
//	Brand       brand.Brand `json:"brand"`
//	Colour      string      `json:"colour"`
//	Description string      `json:"description"`
//	Variants    []struct {
//		Id       string `json:"id"`
//		Title    string `json:"title"`
//		Quantity uint32 `json:"quantity"`
//		Stock    []struct {
//			Id string `json:"id"`
//		} `json:"stock"`
//	} `json:"variants"`
//	Images  []draw.Image `json:"images"`
//	Related []Product    `json:"related"`
//}


type Brand struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type Product struct {
	Id            int       `json:"id" db:"id"`
	ProductLineId int       `json:"-" db:"product_line_id"`
	Name          string    `json:"name" db:"product_line_name"`
	Brand         string    `json:"brand" db:"brand_name"`
	Colour        string    `json:"colour" db:"colour"`
	Description   string    `json:"description" db:"description"`
	Stock         []Stock   `json:"stock"`
	Pictures      []string  `json:"pictures"`
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
