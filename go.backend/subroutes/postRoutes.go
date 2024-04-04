package subroutes

import (
	"context"
	"encoding/json"
	"fmt"
	"gobackend/mongodb"
	"gobackend/mongodb/models"
	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err = godotenv.Load(".env")

type photoPost struct {
	Name   string `json:"name"`
	Prompt string `json:"prompt"`
	Photo  string `json:"photo"`
}

func PostImage(c *gin.Context) {

	coll := mongodb.GetCollection()
	if err != nil {
		log.Fatal("Could not load environment variable in postImage function")
	}
	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))

	var body photoPost
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		http.Error(c.Writer, "Error decoding request body", http.StatusBadRequest)
		return
	}
	resp, err := cld.Upload.Upload(context.TODO(), body.Photo, uploader.UploadParams{})
	if err != nil {
		http.Error(c.Writer, "Error uploading image", http.StatusBadRequest)
		return
	}
	var newPost = photoPost{
		Name:   body.Name,
		Prompt: body.Prompt,
		Photo:  resp.SecureURL,
	}
	cursor, err := coll.InsertOne(context.TODO(), newPost)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(cursor.InsertedID)
	return
}

func GetImages(c *gin.Context) {
	coll := mongodb.GetCollection()
	opts := options.Find()
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		fmt.Print("No Documents found.")
	}
	var results []models.Post
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// for _, result := range results {
	// 	result, _ = bson.MarshalExtJSON(result, false, false)
	// 	_=res
	// }
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(results)
	return
}
