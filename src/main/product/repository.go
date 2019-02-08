package product

import (
	"../models"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	_"log"
)

type Repository struct {
	Connection *sql.DB
}

func NewProductRepository(db *sql.DB) *Repository {
	return &Repository{db }
}

func (r *Repository) Brands() []models.Brand {
	query := "SELECT * FROM brand"

	rows, err := r.Connection.Query(query)
	if err != nil {
		return nil
	}

	var brands []models.Brand

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

func (r *Repository) Search(term string) ([]models.Search, error) {
	query := `SELECT id, name FROM product_line WHERE name LIKE '%' || ? || '%';`

	rows, err := r.Connection.Query(query, term)
	
	if err != nil {
		return nil, err
	}

	var searches []models.Search

	for rows.Next() {
		var search models.Search
		
		err = rows.Scan(
			&search.Id, 
			&search.Name, 
		)

		if err != nil {
			return nil, err
		}

		searches = append(searches, search)
	}

	return searches, nil
}

func (r *Repository) related(productID string) ([]models.Related, error) {
	var relatedList []models.Related
	
	var productLineId int
	
	query := "SELECT product_line_id FROM product WHERE id LIKE ?;"

	r.Connection.QueryRow(query, productID).Scan(&productLineId)
	
	rows, err := r.Connection.Query(`
		SELECT id, colour FROM product 
			WHERE product_line_id LIKE ? AND id NOT LIKE ?;
		`, productLineId, productID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var related models.Related
		err = rows.Scan(&related.Id, &related.Title)

		if err != nil {
			return nil, err
		}

		relatedList = append(relatedList, related)
	}

	return relatedList, nil
}

func (r *Repository) stock(productID string) ([]models.Stock, error) {
	var sizes []models.Stock
	query := `SELECT size, quantity FROM stock
				WHERE stock.product_id like ? ORDER BY ABS(size) ASC;`

	rows, err := r.Connection.Query(query, productID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var size models.Stock
		err = rows.Scan(&size.Size, &size.Quantity)

		if err != nil {
			return nil, err
		}

		sizes = append(sizes, size)
	}

	return sizes, nil
}

func (r *Repository) product(productID string) (models.Product, error) {
	query := `SELECT product.id, product_line.name, product.colour, brand.name AS brand_name, product_line.description FROM product
				INNER JOIN product_line
					ON product.product_line_id = product_line.id 
				INNER JOIN brand 
					ON product_line.brand_id = brand.id 
				WHERE product.id LIKE ? 
				LIMIT 1;`
				
	rows, err := r.Connection.Query(query, productID)
	var product models.Product

	if err != nil {
		return product, err
	}

	for rows.Next() {
		product, err = r.parseProduct(rows)

		if err != nil {
			return product, err
		}
	}

	related, err := r.related(productID)

	if err != nil {
		return product, err
	}

	stock, err := r.stock(productID)

	if err != nil {
		return product, err
	}

	product.Stock = stock
	product.Related = related

	return product, nil
}

func (r *Repository) parseProduct(row *sql.Rows) (models.Product, error) {
	var product models.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Colour,
		&product.Brand,
		&product.Description,
	)

	if err != nil {
		return product, err
	}

	return product, nil
}