package controllers

import (
	"context"
	"encoding/json"
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

var collectionUser *mongo.Collection = database.GetCollection(database.MONGO, "users")

func ShowUsersMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.UserMongo
	defer cancel()
	cur, err := collectionUser.Find(ctx, bson.D{{}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting users: " + err.Error(),
		})
		return
	}
	log.Println("cur: ", cur)
	if err := cur.All(ctx, &users); err != nil {
		log.Fatal("error: ", err)
	}
	c.JSON(200, users)
}

func ShowUserMongoByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserMongo
	defer cancel()
	id := c.Param("id")
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	err = collectionUser.FindOne(ctx, bson.M{"id": newid}).Decode(&user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}
	teste, _ := json.Marshal(user)
	err = json.Unmarshal(teste, &user)
	fmt.Println(user.Name)
	fmt.Println(user.Password)
	c.JSON(200, user)
}

func CreateUserMongo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserMongo
	defer cancel()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	newUser := models.UserMongo{
		ID:        primitive.NewObjectID(),
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := collectionUser.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal("error: ", err)
	}
	id := res.InsertedID
	log.Println("inserted id: ", id)
	c.JSON(200, newUser)
}

func UpdateUserMongoByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	var user models.UserMongo
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	_, err = collectionUser.UpdateOne(ctx, bson.M{"id": newid}, bson.M{"$set": bson.M{
		"name":        user.Name,
		"description": user.Password,
		"updatedat":   time.Now(),
	}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error updating user",
		})
		return
	}
	c.JSON(200, user)
}
func DeleteUserMongoByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	newid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id must be a valid hex",
		})
		return
	}
	_, err = collectionUser.DeleteOne(ctx, bson.M{"id": newid})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error deleting user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}
