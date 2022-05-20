package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		games := main.Group("games")
		{
			games.GET("/:id", controllers.ShowGame)
			games.GET("/", controllers.ShowGames)
			games.POST("/", controllers.CreateGame)
			games.PUT("/:id", controllers.UpdateGame)
			games.DELETE("/:id", controllers.DeleteGame)
		}
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
	}
	return router
}
