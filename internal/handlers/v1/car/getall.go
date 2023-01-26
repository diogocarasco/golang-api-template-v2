package handlers

import (
	"net/http"

	"github.com/diogocarasco/golang-api-template-v2/internal/models/car"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	cars, err := car.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	if len(cars) == 0 {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, cars)
}
