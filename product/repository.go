package product

import (
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) pictures(productID string) ([]string, error) {
	var pictures []string

	query := `SELECT product_image.url FROM product_image 
				WHERE product_image.product_id LIKE ?`

	if err := r.Database.Get(&pictures, query, productID); err != nil {
		return nil, err
	}

	return pictures, nil
}

func (r *Repository) product(productID string) (ContainerProduct, error) {
	var products ContainerProduct
	query := `SELECT product.id, product.name, brand.name AS brand_name, product.description, product.item_quantity, product.item_quantity_postfix FROM product
				INNER JOIN brand 
					ON product.brand_id = brand.id 
				WHERE product.id LIKE ?;`

	if err := r.Database.GetOne(productID, &products, query); err != nil {
		return ContainerProduct{}, err
	}

	return products, nil
}
