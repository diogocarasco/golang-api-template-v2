package handlers

import (
	"net/http"

	"github.com/diogocarasco/golang-api-template-v2/internal/models/car"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Get(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	car, err := car.Get(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if car.Model == "" {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, car)
	}

}
