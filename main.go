package main

import (
	"fmt"
	"go-api/src/models"
	"go-api/src/services"
	"go-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	db := models.OpenDatabaseConnection()
	srv := services.NewUserService(db)
	user, err := srv.GetFirstUsers()
	if err != nil {
		panic("Error fetching first user" + err.Error())
	} else {
		fmt.Println("first user fetched successfully", user.ID)
	}

	router := gin.Default()
	router.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	router.Run(":8080")
}
