package books

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	CollectionBooks = "books"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/books")
	routes.POST("/", AddBook)
	routes.GET("/", GetBooks)
	routes.GET("/:id", GetBook)
	routes.PUT("/:id", UpdateBook)
	routes.DELETE("/:id", DeleteBook)
	log.Println("Routes registered successfully")
}

func getCollection(collectionName string) *mongo.Collection {
	return configs.GetCollection(collectionName)
}
