package books

import (
	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/go-gin-api/pkg/common/models"
	"net/http"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book
	if result := h.DB.Find(&books); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &books)
}
