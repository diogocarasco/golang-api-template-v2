package handlers

import (
	"net/http"

	"github.com/diogocarasco/golang-api-template-v2/internal/models/car"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	row := car.Car{}

	if err := c.ShouldBindJSON(&row); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := car.Update(row)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusOK, id)
	}
}
