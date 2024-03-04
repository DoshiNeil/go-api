package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"uniqueIndex" json:"firstName"`
	LastName  string `gorm:"uniqueIndex" json:"lastName"`
	Email     string `gorm:"not null" json:"email"`
	Country   string `gorm:"not null" json:"country"`
	Age       int    `gorm:"not null;size:3" json:"age"`
}
