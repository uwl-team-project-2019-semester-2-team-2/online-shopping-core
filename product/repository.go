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

	if err := r.Database.Get(productLineId, &relatedList, query); err != nil {
		return nil, err
	}

	return relatedList, nil
}

func (r *Repository) stock(productID string) ([]model.Stock, error) {
	var sizes []model.Stock
	query := `SELECT size, quantity FROM stock WHERE stock.product_id like ? ORDER BY ABS(size) ASC;`

	if err := r.Database.Get(productID, &sizes, query); err != nil {
		return nil, err
	}

	return sizes, nil
}

func (r *Repository) product(productID string) (model.Product, error) {
	var products []model.Product
	query := `SELECT product.id, product_line.id AS product_line_id, product_line.name AS product_line_name,
				product.colour, brand.name AS brand_name, product_line.description FROM product
				INNER JOIN product_line
					ON product.product_line_id = product_line.id 
				INNER JOIN brand 
					ON product_line.brand_id = brand.id 
				WHERE product.id LIKE ? 
				LIMIT 1;`

	if err := r.Database.Get(productID, &products, query); err != nil {
		return model.Product{}, err
	}

	return products[0], nil
}
