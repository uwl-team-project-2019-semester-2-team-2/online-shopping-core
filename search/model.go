package search

type URLParams struct {
	Filter	string   `schema:"filter"`
	Order string `schema:"order"`
	Page int `schema:"page"`
}

type PageInfo struct {
	Page                    int                 `json:"page"`
	Order                   string              `json:"order"`
}

type DietaryFilter struct {
	Name                    string      `json:"name" db:"name"`
	URL                     string      `json:"url" db:"url"`
}

type DatabaseContainer struct {
	Id     					int    	    `json:"id" db:"id"`
	Name    				string 	    `json:"name" db:"name"`
	Price 					string 	    `json:"price" db:"price"`
	CoverPhoto 				string 	    `json:"cover_photo_url" db:"url"`
	ItemQuantity			int  	    `json:"item_quantity" db:"item_quantity"`
	ItemQuantityPostFix 	string      `json:"item_quantity_postfix" db:"item_quantity_postfix"`
}

type Marshaller struct {
	PageInfo                PageInfo            `json:"page_info"`
	Count                   int                 `json:"count"`
	SearchProducts          []DatabaseContainer `json:"products"`
	Filters					[]DietaryFilter     `json:"filters"`
}

type PackInfo struct {
	Quantity     int        `json:"quantity"`
	Postfix      string     `json:"postfix"`
}