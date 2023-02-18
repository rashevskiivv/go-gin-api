package books

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type AddBookRequestBody struct {
	Id          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Description string             `json:"description"`
}

func AddBook(c *gin.Context) {
	body := AddBookRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	var book models.Book
	book.Id = primitive.NewObjectID()
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	result, err := getCollection(CollectionBooks).InsertOne(c, &book)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
