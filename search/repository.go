package search

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) count(term string) (int, error) {
	var quantity int

	query := `SELECT COUNT(*) FROM product
				WHERE product.name LIKE ?`

	if err := r.Database.Count("%"+term+"%", &quantity, query); err != nil {
		return 0, err
	}
	return quantity, nil
}

func (r *Repository) filters() ([]DietaryFilter, error) {
	var filters []DietaryFilter

	query := `SELECT dietary.name, dietary.url FROM dietary`

	if err := r.Database.GetSlice(&filters, query); err != nil {
		return nil, err
	}

	return filters, nil
}

func (r *Repository) search(term string, page int, order string, filters ...string) ([]DatabaseContainer, error) {
	perPage := 25
	upperRange := page * perPage
	lowerRange := upperRange - perPage

	var searches []DatabaseContainer
	var query string
	var orderQuery string
	var filtersQuery string

	switch order {
	case "ascending":
		orderQuery = "ORDER BY product.price ASC"
	case "descending":
		orderQuery = "ORDER BY product.price DESC"
	default:
		orderQuery = "ORDER BY product.id DESC"
	}

	if filters != nil {
		filtersQuery = `AND product.id NOT IN (SELECT product_dietary.product_id
						FROM product_dietary
						JOIN dietary ON product_dietary.dietary_id = dietary.id
						WHERE dietary.url IN (:filters))`
	}

	query = fmt.Sprintf(`SELECT
					product.id,
					product.name,
					product.price,
					product.item_quantity,
					product.item_quantity_postfix,
					product_image.url
				FROM product
    			JOIN product_image_cover ON product.id = product_image_cover.product_id
    			JOIN product_image ON product_image_cover.product_image_id = product_image.id
				WHERE product.name LIKE :term 
				%s
				%s 
    			LIMIT %d, %d;`, filtersQuery, orderQuery, lowerRange, upperRange)

	queryMap := map[string]interface{}{
		"term":  "%"+term+"%",
		"filters": filters,
	}

	query, args, err := sqlx.Named(query, queryMap)
	query, args, err = sqlx.In(query, args...)
	query = r.Database.Connection.Rebind(query)
	rows, err := r.Database.Connection.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var search DatabaseContainer
		err := rows.StructScan(&search)

		searches = append(searches, search)
		if err != nil {
			return nil, err
		}
	}

	return searches, nil
}
