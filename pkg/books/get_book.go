package books

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}

	err = getCollection(CollectionBooks).FindOne(c, bson.M{"id": objId}).Decode(&book)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, book)
}
