package online_shopping_core

import (
	"errors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/product"
	"log"
)

func StartServer() (*typhon.Server, error) {
	_, err := database.OpenMysql("root", "liam", "127.0.0.1", "shop")

	if err != nil {
		return nil, errors.New("unable to start server: " + err.Error())
	}

	svc := getRouter().Serve().
		Filter(typhon.ErrorFilter)
	srv, err := typhon.Listen(svc, ":8080")
	if err != nil {
		return nil, errors.New("unable to start server: " + err.Error())
	}
	log.Printf("Server listening on %s...", srv.Listener().Addr())

	return srv, nil
}

func getRouter() *typhon.Router {
	router := &typhon.Router{}
	product.Routes(router)

	return router
}
