package main

import (
	"fmt"
	"go-api/src/models"
	"go-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	db := models.OpenDatabaseConnection()

	newUser := models.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@gmail.com",
		Country:   "Spain",
		Age:       25,
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create a new user" + result.Error.Error())
	}

	fmt.Printf("New User added successfully")

	router := gin.Default()
	router.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	router.Run(":8080")
}
