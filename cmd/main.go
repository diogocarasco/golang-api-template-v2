package main

import (
	"fmt"

	"github.com/diogocarasco/golang-api-template-v2/config"
	carHandler "github.com/diogocarasco/golang-api-template-v2/internal/handlers/v1/car"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.ForceConsoleColor()
	r := gin.Default()

	if err := config.Load(); err != nil {
		fmt.Println("Falha ao carregar as configurações")
	}

	/*
		db, err := db.OpenConnection()
		if err != nil {
			fmt.Println("Falha ao conectar no banco de dados")
		} else {
			fmt.Println("Migrando base de dados")
			db.AutoMigrate(&car.Car{})
		}
	*/

	v1 := r.Group("/v1")
	{
		v1.GET("/car/", carHandler.GetAll)
		v1.GET("/car/:id", carHandler.Get)
		v1.POST("/car", carHandler.Create)
		v1.PUT("/car/:id", carHandler.Update)
		v1.DELETE("/car/:id", carHandler.Delete)
	}

	r.Run(":" + config.GetAppPort())

}
