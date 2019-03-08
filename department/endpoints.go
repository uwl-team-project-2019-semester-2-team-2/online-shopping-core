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

	dept, err := pr.Repository.departments()

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	return r.Response(dept)
}
