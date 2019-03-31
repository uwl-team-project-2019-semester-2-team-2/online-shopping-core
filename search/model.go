package search

type URLParams struct {
	Filter string `schema:"filter"`
	Order  string `schema:"order"`
	Page   int    `schema:"page"`
}

type UserFilters struct {
	Inclusive []string
	Exclusive []string
}

type PageInfo struct {
	Page          int      `json:"page"`
	Order         string   `json:"order"`
	Count         int      `json:"product_count"`
	ActiveFilters []string `json:"active_filters"`
}

type DietaryFilter struct {
	Id       string `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	URL      string `json:"url" db:"url"`
	Filter   bool   `json:"filter" db:"filter"`
	Quantity int    `json:"quantity" db:"-"`
}

type DatabaseContainer struct {
	Id                  int    `json:"id" db:"id"`
	Name                string `json:"name" db:"name"`
	Price               string `json:"price" db:"price"`
	DepartmentId        string `json:"department_id" db:"department_id"`
	ProductDepartment   string `json:"department_name" db:"department_name"`
	CoverPhoto          string `json:"cover_photo_url" db:"url"`
	ItemQuantity        int    `json:"item_quantity" db:"item_quantity"`
	ItemQuantityPostFix string `json:"item_quantity_postfix" db:"item_quantity_postfix"`
}

type Marshaller struct {
	PageInfo       PageInfo            `json:"page_info"`
	SearchProducts []DatabaseContainer `json:"products"`
	Filters        []DietaryFilter     `json:"filters"`
	Department     []Department        `json:"departments"`
}

type ProductDepartment struct {
	Id   int    `json:"id" db:"database_id"`
	Name string `json:"name" db:"database_name"`
}

type Department struct {
	Id       int    `json:"id" db:"db"`
	Name     string `json:"name" db:"name"`
	URL      string `json:"url" db:"url"`
	Quantity int    `json:"quantity" db:"quantity"`
}

type Brand struct {
}

type PackInfo struct {
	Quantity int    `json:"quantity"`
	Postfix  string `json:"postfix"`
}
