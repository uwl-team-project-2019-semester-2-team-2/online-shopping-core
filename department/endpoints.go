package department

import (
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
)

type Department struct {
	Repository *Repository
}

func Init(database database.Database, router *typhon.Router) {
	repo := Repository{&database}
	prod := Department{&repo}
	prod.Routes(router)
}

func (pr *Department) Routes(router *typhon.Router) {
	router.GET("/department", pr.List)
}

func (pr *Department) List(r typhon.Request) typhon.Response {
	response := typhon.NewResponse(r)

	depts, err := pr.Repository.departments()

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	var sortedDepts []Container

	for _, dept := range depts {
		if dept.ParentId == 0 {
			parentDept := Container {
				Id: dept.Id,
				Name: dept.Name,
				URL: dept.URL,
				Children: buildDepartment(dept.Id, depts),
			}

			sortedDepts = append(sortedDepts, parentDept)
		}
	}

	return r.Response(sortedDepts)
}

func buildDepartment(id int, departments []Container) []Container {

	var childDepts []Container

	for _, dept := range departments {
		if dept.ParentId == id {
			childDept := Container {
				Id: dept.Id,
				Name: dept.Name,
				URL: dept.URL,
				Children: buildDepartment(dept.Id, departments),
			}
			childDepts = append(childDepts, childDept)
		}
	}

	return childDepts
}