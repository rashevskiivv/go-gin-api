package books

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	result, err := getCollection(CollectionBooks).DeleteOne(c, bson.M{"id": objId})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	if result.DeletedCount < 1 {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("user with specified ID not found"))
		log.Println(err)
		return
	}
	c.Status(http.StatusOK)
}
