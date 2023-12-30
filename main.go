package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
)

func main() {
	//Initialize Database
	database.Connect("root:@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()

	//Initialize Router
	router := initRouter()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running at port :8080")
		return
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Entry)
	api := router.Group("/api")
	{
		api.POST("/token", controllers.TokenBuilder)
		api.POST("/user/register", controllers.UserRegister)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
