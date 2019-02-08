package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/services"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = "root"
	mysqlConfig.Passwd = "liam"
	mysqlConfig.Addr = "127.0.0.1"
	mysqlConfig.DBName = "shop"

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	router := typhon.Router{}
	router.GET("/product/:productId", services.GetProduct)

	startServer(router)
}

func startServer(router typhon.Router) {
	svc := router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)
	srv, err := typhon.Listen(svc, ":8000")
	if err != nil {
		panic(err)
	}
	log.Printf("Server listening on %s...", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("Interrupt or terminate received, shutting down...")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
