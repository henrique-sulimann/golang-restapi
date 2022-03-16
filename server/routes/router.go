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
			gamesMongo.GET("/:id", controllers.ShowGameMongo)
			gamesMongo.GET("/", controllers.ShowGamesMongo)
			gamesMongo.POST("/", controllers.CreateGameMongo)
			gamesMongo.PUT("/:id", controllers.UpdateGameMongo)
			gamesMongo.DELETE("/:id", controllers.DeleteGameMongo)
		}
	}
	return router
}
