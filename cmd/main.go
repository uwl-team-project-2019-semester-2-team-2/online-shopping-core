package main

import (
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv, err := online_shopping_core.StartServer()

	if err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("Interrupt or terminate received, shutting down...")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}