package product

type Brand struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type Picture struct {
	URL string `json:"url" db:"url"`
}

type PackInfo struct {
	Quantity int    `json:"quantity"`
	Postfix  string `json:"postfix"`
}

type ContainerProduct struct {
	Id           int      `json:"id" db:"id"`
	Name         string   `json:"name" db:"name"`
	Brand        string   `json:"brand" db:"brand_name"`
	Description  string   `json:"description" db:"description"`
	ItemQuantity int      `json:"-" db:"item_quantity"`
	Postfix      string   `json:"-" db:"item_quantity_postfix"`
	Pictures     []string `json:"pictures"`
	PackInfo     PackInfo `json:"package_info"`
}
