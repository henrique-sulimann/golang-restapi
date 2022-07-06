package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Teste(c *gin.Context) {
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "unauthorized",
		})
		return
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "unauthorized",
		})
		return
	}
	fmt.Println(userId)
	c.JSON(200, "teste")
}
