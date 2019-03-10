package product

import (
	"fmt"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/model"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) related(productId string, productLineId string) ([]model.Related, error) {
	var relatedList []model.Related
	query := fmt.Sprintf(`SELECT id, colour FROM product WHERE product_line_id LIKE ? AND id NOT LIKE "%s";`, productId)

	if err := r.Database.Get(&relatedList, query, productLineId); err != nil {
		return nil, err
	}

	return relatedList, nil
}

func (r *Repository) stock(productID string) ([]model.Stock, error) {
	var sizes []model.Stock
	query := `SELECT size, quantity FROM stock WHERE stock.product_id like ? ORDER BY ABS(size) ASC;`

	if err := r.Database.Get(&sizes, query, productID); err != nil {
		return nil, err
	}

	return sizes, nil
}

func (r *Repository) pictures(productID string) ([]model.Picture, error) {
	var pictures []model.Picture

	query := `SELECT product_image.url FROM product_image 
				WHERE product_image.product_id LIKE ?`

	if err := r.Database.Get(&pictures, query, productID); err != nil {
		return nil, err
	}

	return pictures, nil
}

func (r *Repository) product(productID string) (model.Product, error) {
	var products model.Product
	query := `SELECT product.id, product.name, brand.name AS brand_name, product.description FROM product
				INNER JOIN brand 
					ON product.brand_id = brand.id 
				WHERE product.id LIKE ?;`

	if err := r.Database.GetOne(productID, &products, query); err != nil {
		return model.Product{}, err
	}

	return products, nil
}
