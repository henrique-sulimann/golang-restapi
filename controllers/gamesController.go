package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi/database"
	"github.com/henrique-sulimann/golang-restapi/models"
)

func ShowGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a number",
		})
		return
	}
	db := database.GetDatabase()
	var game models.Game
	err = db.First(&game, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "game not found",
		})
		return
	}
	c.JSON(200, game)
}

func CreateGame(c *gin.Context) {
	db := database.GetDatabase()
	var game models.Game
	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	err = db.Create(&game).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error creating game",
		})
		return
	}
	c.JSON(200, game)
}

func ShowGames(c *gin.Context) {
	db := database.GetDatabase()
	var games []models.Game
	err := db.Find(&games).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting games",
		})
		return
	}
	c.JSON(200, games)
}

func UpdateGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a number",
		})
		return
	}
	db := database.GetDatabase()
	var game models.Game
	err = db.First(&game, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "game not found",
		})
		return
	}
	err = c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	err = db.Save(&game).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error updating game",
		})
		return
	}
	c.JSON(200, game)
}

func DeleteGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a number",
		})
		return
	}
	db := database.GetDatabase()
	var game models.Game
	err = db.First(&game, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "game not found",
		})
		return
	}
	err = db.Delete(&game).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error deleting game",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "game deleted",
	})
}
