package books

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type UpdateBookRequestBody struct {
	Id          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Description string             `json:"description"`
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	body := UpdateBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	update := bson.M{"title": body.Title, "author": body.Author, "description": body.Description}
	result, err := getCollection(CollectionBooks).UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	var book models.Book
	if result.MatchedCount == 1 {
		err = getCollection(CollectionBooks).FindOne(c, bson.M{"id": objId}).Decode(&book)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			log.Println(err)
			return
		}
	}
	c.JSON(http.StatusOK, &book)
}
