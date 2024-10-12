package main

import (
	"project/rest_api/db"
	"project/rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main()  {
	db.InitDb()
	server:=gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
