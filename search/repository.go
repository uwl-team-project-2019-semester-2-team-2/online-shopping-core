package search

import (
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/model"
	"strconv"
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

func (r *Repository) search(term string, page int) ([]model.Search, error) {
	perPage := 25
	upperRange := page * perPage
	lowerRange := strconv.Itoa(upperRange - perPage)

	var searches []model.Search
	query := `SELECT product.id, product.name, product.price, product.item_quantity, product.item_quantity_postfix FROM product
				WHERE product.name LIKE ?
				ORDER BY product.id
    			LIMIT ` + lowerRange + `, ` + strconv.Itoa(upperRange) + `;`

	if err := r.Database.Get("%"+term+"%", &searches, query); err != nil {
		return nil, err
	}

	return searches, nil
}
