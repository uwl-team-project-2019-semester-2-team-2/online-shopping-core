package product

import (
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/brand"
	"image/draw"
)

type Product struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Brand       brand.Brand `json:"brand"`
	Description string      `json:"description"`
	Variants    []struct {
		Id       string `json:"id"`
		Title    string `json:"title"`
		Colour   string `json:"colour"`
		Quantity uint32 `json:"quantity"`
		Stock    []struct {
			Id string `json:"id"`
		} `json:"stock"`
	} `json:"variants"`
	Images  []draw.Image `json:"images"`
	Related []Product    `json:"related"`
}
