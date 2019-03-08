package model

type Search struct {
	Id     					int    	`json:"id" db:"id"`
	Name    				string 	`json:"name" db:"name"`
	Variant 				string 	`json:"variant" db:"colour"`
	Price 					string 	`json:"price" db:"price"`
	ItemQuantity			int  	`json:"item_quantity" db:"item_quantity"`
	ItemQuantityPostFix 	string  `json:"item_quantity_postfix" db:"item_quantity_postfix"`
}
