package online_shopping_core

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/department"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/product"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/search"
	"log"
)

func StartServer() (*typhon.Server, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = "test"
	mysqlConfig.Passwd = "test"
	mysqlConfig.Addr = "127.0.0.1"
	mysqlConfig.DBName = "shop"

	db, err := database.Connect(mysqlConfig)

	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	router := initModules(db)
	svc := router.Serve().Filter(typhon.ErrorFilter)

	srv, err := typhon.Listen(svc, ":8080")
	if err != nil {
		return nil, errors.New("unable to start server: " + err.Error())
	}

	log.Printf("Server listening on %s...", srv.Listener().Addr())

	return srv, nil
}

func initModules(db database.Database) *typhon.Router {
	router := &typhon.Router{}
	product.Init(db, router)
	search.Init(db, router)
	department.Init(db, router)
	return router
}
