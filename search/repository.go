package search

import (
	"fmt"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) count(term string, filters UserFilters) (int, error) {
	var exclusiveQuery string
	var inclusiveQuery string

	queryMap := map[string]interface{}{
		"term":      "%" + term + "%",
		"inclusive": filters.Inclusive,
		"exclusive": filters.Exclusive,
	}

	if filters.Exclusive != nil {
		exclusiveQuery = `AND product.id NOT IN (SELECT product_dietary.product_id
						FROM product_dietary
						JOIN dietary ON product_dietary.dietary_id = dietary.id
						WHERE dietary.url IN (:exclusive))`
	}

	if filters.Inclusive != nil {
		inclusiveQuery = `AND product.id IN (SELECT product_dietary.product_id
						FROM product_dietary
						JOIN dietary ON product_dietary.dietary_id = dietary.id
						WHERE dietary.url IN (:inclusive))`
	}

	query := fmt.Sprintf(`SELECT COUNT(*) FROM product
				WHERE product.name LIKE :term
				%s %s`, inclusiveQuery, exclusiveQuery)

	rows, err := r.Database.Query(query, queryMap)

	if err != nil {
		return 0, err
	}

	var quantity int
	err = r.Database.Scan(&quantity, rows)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (r *Repository) filters() ([]DietaryFilter, error) {
	var filters []DietaryFilter

	query := `SELECT dietary.id, dietary.name, dietary.url, dietary.filter FROM dietary`

	if err := r.Database.GetSlice(&filters, query); err != nil {
		return nil, err
	}

	return filters, nil
}

func (r *Repository) filterCount(term string) ([]DietaryFilter, error) {
	var outFilters []DietaryFilter

	filters, err := r.filters()

	if err != nil {
		return nil, err
	}

	for _, filter := range filters {

		queryMap := map[string]interface{}{
			"term": "%" + term + "%",
			"id":   filter.Id,
		}

		var filterQ string
		if filter.Filter {
			filterQ = "!="
		} else {
			filterQ = "="
		}

		filterQuery := fmt.Sprintf(`SELECT COUNT(*)
				FROM product
				JOIN product_dietary ON product_dietary.product_id = product.id
    			WHERE product_dietary.dietary_id %s :id
				AND product.name LIKE :term;`, filterQ)

		rows, err := r.Database.Query(filterQuery, queryMap)
		err = r.Database.Scan(&filter.Quantity, rows)

		outFilters = append(outFilters, filter)

		if err != nil {
			return nil, err
		}
	}

	return outFilters, nil
}

func (r *Repository) search(term string, page int, order string, filters UserFilters) ([]DatabaseContainer, error) {
	perPage := 25
	upperRange := page * perPage
	lowerRange := upperRange - perPage

	var searches []DatabaseContainer
	var query string
	var orderQuery string
	var exclusiveQuery string
	var inclusiveQuery string

	switch order {
	case "ascending":
		orderQuery = "ORDER BY product.price ASC"
	case "descending":
		orderQuery = "ORDER BY product.price DESC"
	default:
		orderQuery = "ORDER BY product.id DESC"
	}

	if filters.Exclusive != nil {
		exclusiveQuery = `AND product.id NOT IN (SELECT product_dietary.product_id
						FROM product_dietary
						JOIN dietary ON product_dietary.dietary_id = dietary.id
						WHERE dietary.url IN (:exclusive))`
	}

	if filters.Inclusive != nil {
		inclusiveQuery = `AND product.id IN (SELECT product_dietary.product_id
						FROM product_dietary
						JOIN dietary ON product_dietary.dietary_id = dietary.id
						WHERE dietary.url IN (:inclusive))`
	}

	query = fmt.Sprintf(`SELECT
					product.id,
					product.name,
					product.price,
					department.name as department_name,
					department.id as department_id,
					product.item_quantity,
					product.item_quantity_postfix,
					product_image.url
				FROM product
				JOIN department ON product.department_id = department.id
    			JOIN product_image_cover ON product.id = product_image_cover.product_id
    			JOIN product_image ON product_image_cover.product_image_id = product_image.id
				WHERE product.name LIKE :term 
				%s %s %s 
    			LIMIT %d, %d;`, inclusiveQuery, exclusiveQuery, orderQuery, lowerRange, upperRange)

	queryMap := map[string]interface{}{
		"term":      "%" + term + "%",
		"exclusive": filters.Exclusive,
		"inclusive": filters.Inclusive,
	}

	rows, err := r.Database.Query(query, queryMap)

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
