package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Description string             `json:"description"`
}
