package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		gamesMongo := main.Group("gamesmongo")
		{
			gamesMongo.GET("/:id", controllers.ShowGameMongoByID)
			gamesMongo.GET("/", controllers.ShowGamesMongo)
			gamesMongo.POST("/", controllers.CreateGameMongo)
			gamesMongo.PUT("/:id", controllers.UpdateGameMongoByID)
			gamesMongo.DELETE("/:id", controllers.DeleteGameMongo)
		}
		usersMongo := main.Group("usersmongo")
		{
			usersMongo.GET("/:id", controllers.ShowUserMongoByID)
			usersMongo.GET("/", controllers.ShowUsersMongo)
			usersMongo.POST("/", controllers.CreateUserMongo)
			usersMongo.PUT("/:id", controllers.UpdateUserMongoByID)
			usersMongo.DELETE("/:id", controllers.DeleteUserMongoByID)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		exemplo := main.Group("teste")
		{
			exemplo.GET("/", fmt.Println("teste"))
		}
	}
	return router
}
