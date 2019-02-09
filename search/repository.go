package search

import (
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/model"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) search(term string) ([]model.Search, error) {
	var searches []model.Search
	query := `SELECT product.id, product_line.name, product.colour FROM product 
				INNER JOIN product_line
					ON product.product_line_id = product_line.id
				WHERE product_line.name LIKE ?
				ORDER BY product.id
    			LIMIT 0,25;`

	if err := r.Database.Get("%" + term + "%", &searches, query); err != nil {
		return nil, err
	}

	return searches, nil
}