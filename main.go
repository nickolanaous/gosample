package main

import (
	"github.com/gin-gonic/gin"
	"go.naous.net/api/db"
	"go.naous.net/api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":1337")

}
