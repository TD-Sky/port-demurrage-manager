package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"log"
	"service/dba"
	"service/middleware"
	"service/routers"
	"service/utils"
)

func main() {
	db, err := sqlx.Connect("postgres",
		"user=postgres dbname=port-demurrage-db sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		log.Fatalln(err)
	}

	server := gin.Default()
	server.Use(middleware.StateContext(db))

	routers.General(server)

	cron := cron.New(cron.WithLocation(utils.CN))
	cron.AddFunc("@daily", dba.Demurrage_updater(db)) // 计费
	cron.AddFunc("0 8 * * *", dba.Shipment(db))       // 开船
	cron.Start()

	server.Run()
}
