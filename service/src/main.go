package main

import (
	"log"
	"service/routers"
	"service/state"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=port-demurrage-db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	server := gin.Default()
	state := state.State{Database: db}

	routers.General_router(server, &state)

	server.Run()
}
