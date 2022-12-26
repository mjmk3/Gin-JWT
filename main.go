package main

import (
	"Gin-JWT/controller"
	"Gin-JWT/db"
	"Gin-JWT/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	db.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true")
	db.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":5500")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
