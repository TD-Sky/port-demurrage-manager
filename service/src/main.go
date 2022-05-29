package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"service/middleware"
	"service/routers"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=port-demurrage-db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	server := gin.Default()
	server.Use(middleware.StateContext(db))

	routers.General_router(server)

	server.Run()
}
