package services

import (
	"go-api/src/errors"
	"go-api/src/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetFirstUsers() (*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUser(user *models.User) (bool, error)
	DeleteUser(id int) (bool, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (us *userService) GetFirstUsers() (*models.User, error) {
	var user models.User
	result := us.db.First(&user)
	if result.Error != nil {
		return nil, result.Error

	}
	return &user, nil
}

func (us *userService) GetUserById(id int) (*models.User, error) {
	var user models.User
	result := us.db.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error

	}
	return &user, nil
}

func (us *userService) UpdateUser(updatedUser *models.User) (bool, error) {
	var user models.User
	result := us.db.Where("ID = ?", updatedUser.ID).First(&user)
	if result.Error != nil {
		return false, result.Error

	}

	user.Age = updatedUser.Age
	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	user.Country = updatedUser.Country

	result = us.db.Save(&user)
	if result.Error != nil {
		return false, result.Error

	}

	return true, nil
}

func (us *userService) DeleteUser(id int) (bool, error) {
	user, error := us.GetUserById(id)
	if error != nil {
		return false, error
	}

	result := us.db.Delete(&user)
	if result.Error != nil {
		return false, error
	} else if result.RowsAffected == 0 {
		return false, &errors.NoRowsAffectedError{}
	} else {
		return true, nil
	}
}
