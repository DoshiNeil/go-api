package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

func OpenDatabaseConnection() {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=required TimeZone=Asia/Kolkata", host , username, password, database, port);

	Database, err = gorm.Open(.Open(dsn), )
}
