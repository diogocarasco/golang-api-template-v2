package handlers

import (
	"net/http"

	"github.com/diogocarasco/golang-api-template-v2/internal/models/car"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	rows, err := car.Delete(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, rows)
	}
}
