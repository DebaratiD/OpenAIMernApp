package main

import (
	"gobackend/mongodb"
	"gobackend/subroutes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func init() {
	mongodb.ConnectToDB()
}

func main() {

	//Use := for implicit variable declaration
	router := gin.Default()
	cORS := cors.DefaultConfig()
	cORS.AllowOrigins = []string{"*"}
	// cORS.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	// cORS.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Origin", "Access-Control-Request-Method"}

	router.Use(cors.New(cORS))

	postImage := router.Group("post")
	postImage.POST("/", subroutes.PostImage)
	postImage.GET("/", subroutes.GetImages)

	//Using single quotes instead of double quotes will generate errors.
	router.GET("/", func(ctx *gin.Context) {
		//Add status code with the data to be sent
		ctx.JSON(200, gin.H{
			"data": "Hello",
		})
	})
	//Run function attaches to an HTTP Server and runs on localhost 6000
	router.Run("localhost:8080")
}
