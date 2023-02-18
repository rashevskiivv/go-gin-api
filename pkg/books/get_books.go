package books

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	results, err := getCollection(CollectionBooks).Find(c, bson.M{})
	defer func(results *mongo.Cursor, ctx context.Context) {
		err = results.Close(ctx)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			log.Println(err)
			return
		}
	}(results, c)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		log.Println(err)
		return
	}

	for results.Next(c) {
		var book models.Book
		if err = results.Decode(&book); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			log.Println(err)
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}
