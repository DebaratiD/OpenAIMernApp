package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post represents the structure of a post.
type Post struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Prompt string             `bson:"prompt,omitempty"`
	Photo  string             `bson:"photo,omitempty"`
}

// CreatePostSchema creates a new PostSchema.
func CreatePostSchema() *Post {
	return &Post{}
}
