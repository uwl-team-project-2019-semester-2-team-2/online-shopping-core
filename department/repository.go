package department

import (
	"fmt"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) departments() ([]Container, error) {

	var departmentList []Container
	query := fmt.Sprintf("SELECT id, name, parent_id, url FROM department;")

	if err := r.Database.List(&departmentList, query); err != nil {
		return nil, err
	}

	return departmentList, nil
}
