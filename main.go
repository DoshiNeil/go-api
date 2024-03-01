package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	router.Run(":8080")

	// newUser := User{
	// 	FirstName: "Jane",
	// 	LastName:  "Doe",
	// 	Email:     "jane.doe@gmail.com",
	// 	Country:   "Spain",
	// 	Age:       25,
	// }

}
