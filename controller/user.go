package controller

import (
	"net/http"
	"rest/config"
	"rest/entity"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name    string `json:"name" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Address entity.Address
}

func GetDetails(c *gin.Context) {
	db := config.InitDb()
	var users []entity.User
	db.Preload("Address").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

func CreateUser(c *gin.Context) {
	db := config.InitDb()
	var input CreateUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{Name: input.Name, Gender: input.Gender,
		Age: input.Age, Address: entity.Address{City: input.Address.City,
			Pin: input.Address.Pin, State: input.Address.State}}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
