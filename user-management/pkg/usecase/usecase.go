package usecase

import (
	"errors"
	"strings"
	"user-management/pkg/models"
	"user-management/pkg/repository"

	"github.com/go-playground/validator/v10"
)

type UseCase struct {
	Repo repository.Repositories
}

type Usecase interface {
	CreateUser(models.User) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserById(int) (models.User, error)
	UpdateUserById(int, models.User) error
	DeleteUserById(int) error
}

func NewUseCase(repo repository.Repositories) Usecase {
	return &UseCase{
		Repo: repo,
	}
}

func (us *UseCase) CreateUser(usr models.User) (models.User, error) {
	validate := validator.New()

	err := validate.Struct(usr)
	if err != nil {
		return models.User{}, err
	}

	if !strings.Contains(usr.Email, "@") {
		return models.User{}, errors.New("incorrect email address")
	}

	return us.Repo.CreateUser(usr)
}

func (us *UseCase) GetUsers() ([]models.User, error) {
	return us.Repo.GetUsers()
}

func (us *UseCase) GetUserById(id int) (models.User, error) {
	return us.Repo.GetUserById(id)
}

func (us *UseCase) UpdateUserById(id int, usr models.User) error {
	return us.Repo.UpdateUserById(id, usr)
}

func (us *UseCase) DeleteUserById(id int) error {
	return us.Repo.DeleteUserById(id)
}
