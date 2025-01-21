package main

import (
	"github.com/dsabljic/event-management/db"
	"github.com/dsabljic/event-management/env"
	"github.com/dsabljic/event-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB(env.GetString("DB_ADDR", "postgres://admin:admin@localhost:54332/event-management?sslmode=disable"))
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(env.GetString("ADDR", ":8080"))
}
