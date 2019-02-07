package product

import (
	"../models"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	_"log"
	"fmt"
)

type ProductRepository struct {
	Connection *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository { db }
}

func (pr *ProductRepository) Brands() []models.Brand {
	query := "SELECT * FROM brand"

	rows, err := pr.Connection.Query(query)
	if err != nil {
		return nil
	}

	brands := []models.Brand {}

	for rows.Next() {
		var brand models.Brand
		err = rows.Scan(&brand.Id, &brand.Name, &brand.Description)
		if err != nil {
			return nil
		}

		brands = append(brands, brand)
	}

	return brands
}

func (pr *ProductRepository) GetByID(id string) models.Product {

}

func (pr *ProductRepository) related(productID string) []models.Related {
	var relatedList []models.Related
	
	var prodcutLineId int
	
	query := "SELECT product_line_id FROM product WHERE id LIKE ?;"

	pr.Connection.QueryRow(query, productID).Scan(&prodcutLineId)
	
	rows, err := pr.Connection.Query(`
		SELECT id, colour FROM product 
			WHERE product_line_id LIKE ? AND id NOT LIKE ?;
		`, prodcutLineId, productID)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for rows.Next() {
		var related models.Related
		err = rows.Scan(&related.Id, &related.Title)

		if err != nil {
			return nil
		}

		relatedList = append(relatedList, related)
	}

	return relatedList
}

func (pr *ProductRepository) stock(productID string) []models.Stock {
	var sizes []models.Stock
	query := `SELECT size, quantity FROM stock
				WHERE stock.product_id like ? ORDER BY ABS(size) ASC;`

	rows, err := pr.Connection.Query(query, productID)

	if err != nil {
		return nil
	}

	for rows.Next() {
		var size models.Stock
		err = rows.Scan(&size.Size, &size.Quantity)

		if err != nil {
			return nil
		}

		sizes = append(sizes, size)
	}

	return sizes
}

func (pr *ProductRepository) product(productID string) models.Product {
	query := `SELECT product.id, product_line.name, product.colour, brand.name AS brand_name, product_line.description FROM product
				INNER JOIN product_line
					ON product.product_line_id = product_line.id 
				INNER JOIN brand 
					ON product_line.brand_id = brand.id 
				WHERE product.id LIKE ? 
				LIMIT 1;`
				
	row := pr.Connection.QueryRow(query, productID)

	var product models.Product
	
	row.Scan(
		&product.Id, 
		&product.Name,
		&product.Colour,
		&product.Brand,
		&product.Description,
	)

	product.Stock = pr.stock(productID)
	product.Related = pr.related(productID)
	return product
}

func parseProduct(row *sql.Row) models.Product {



}