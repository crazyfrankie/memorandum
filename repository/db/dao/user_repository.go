package dao

import (
	"errors"

	"gorm.io/gorm"

	"memorandum/repository/db/model"
)

var ErrUserNotFound = errors.New("admin not found")

type UserRepository interface {
	CreateUser(user *model.User) error
	FindByName(name string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

func (u *userRepository) FindByName(name string) (*model.User, error) {
	var user model.User

	result := DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (u *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User

	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
