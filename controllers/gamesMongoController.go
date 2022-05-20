package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi/database"
	"github.com/henrique-sulimann/golang-restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionGames *mongo.Collection = database.GetCollection(database.MONGO, "games")

func ShowGamesMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var games []models.GameMongo
	defer cancel()
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
	cur, err := collectionGames.Find(ctx, bson.D{{}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting games: " + err.Error(),
		})
		return
	}
	log.Println("cur: ", cur)
	if err := cur.All(ctx, &games); err != nil {
		log.Fatal("error: ", err)
	}
	c.JSON(200, games)
}

func CreateGameMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var game models.GameMongo
	defer cancel()
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
	err = c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	newGame := models.GameMongo{
		ID:          primitive.NewObjectID(),
		Name:        game.Name,
		Description: game.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	res, err := collectionGames.InsertOne(ctx, newGame)
	if err != nil {
		log.Fatal("error: ", err)
	}
	id := res.InsertedID
	log.Println("inserted id: ", id)
	c.JSON(200, newGame)
}

func ShowGameMongoByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var game models.GameMongo
	defer cancel()
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
	id := c.Param("id")
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	err = collectionGames.FindOne(ctx, bson.M{"id": newid}).Decode(&game)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "game not found",
		})
		return
	}
	c.JSON(200, game)
}

func UpdateGameMongoByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
	id := c.Param("id")
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	var game models.GameMongo
	err = c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	_, err = collectionGames.UpdateOne(ctx, bson.M{"id": newid}, bson.M{"$set": bson.M{
		"name":        game.Name,
		"description": game.Description,
		"updatedat":   time.Now(),
	}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error updating game",
		})
		return
	}
	c.JSON(200, game)
}
func DeleteGameMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
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
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	_, err = collectionGames.DeleteOne(ctx, bson.M{"id": newid})
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
