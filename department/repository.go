package department

import (
	"fmt"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/model"
)

type Repository struct {
	Database *database.Database
}

func (r *Repository) departments() ([]model.Department, error) {

	var departmentList []model.Department
	query := fmt.Sprintf("SELECT id, name FROM department;")

	if err := r.Database.List(&departmentList, query); err != nil {
		return nil, err
	}

	return departmentList, nil
}
