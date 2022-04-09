package main

import (
	"net/http"
	"rest/config"
	"rest/controller"
	"rest/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := config.InitDb()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "first APi"})
	})
	db.Create(&entity.User{Name: "Rahul", Gender: "Male", Age: 12, Address: entity.Address{City: "pithoragarh", Pin: "123223", State: "UttaraKhand"}})

	router.GET("/users", controller.GetDetails)
	router.POST("/create", controller.CreateUser)
	router.Run()

}
