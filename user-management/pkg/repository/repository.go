package repository

import (
	"errors"
	"strconv"
	"user-management/pkg/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Repositories interface {
	CreateUser(models.User) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserById(int) (models.User, error)
	UpdateUserById(int, models.User) error
	DeleteUserById(int) error
}

func NewRepositories(db *gorm.DB) Repositories {
	return &Repository{
		DB: db,
	}
}

func (re *Repository) CreateUser(usr models.User) (models.User, error) {
	res := re.DB.Create(&usr)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	if res.RowsAffected != 1 {
		return models.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (re *Repository) GetUsers() ([]models.User, error) {
	var users []models.User
	res := re.DB.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil

}

func (re *Repository) GetUserById(id int) (models.User, error) {
	var user models.User

	res := re.DB.First(&user, id)

	if res.Error != nil {
		return models.User{}, res.Error
	}

	if res.RowsAffected != 1 {
		return models.User{}, errors.New("doesn't find the user")
	}

	return user, nil
}
func (re *Repository) UpdateUserById(id int, usr models.User) error {
	data := map[string]interface{}{
		"first_name":   usr.First_Name,
		"last_name":    usr.Last_Name,
		"user_name":    usr.User_Name,
		"email":        usr.Email,
		"phone_number": usr.Phone_Number,
	}
	res := re.DB.Model(&models.User{}).Where("id = ?", id).Updates(data)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.New("failed to update user data")
	}
	return nil
}

func (re *Repository) DeleteUserById(id int) error {
	res := re.DB.Delete(&models.User{}, id)

	if res.RowsAffected != 1 {
		return errors.New("failed to delete" + strconv.Itoa(id) + "this user")
	}

	return nil
}
