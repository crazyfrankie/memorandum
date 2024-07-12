package service

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"memorandum/pkg/util"
	"memorandum/repository/db/dao"
	"memorandum/repository/db/model"
)

type UserService interface {
	RegisterUser(user *model.LoginData) error
	LoginUser(req *model.LoginData) (interface{}, string, error)
}

type userService struct {
	repo dao.UserRepository
}

func NewUserService(repo dao.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) RegisterUser(user *model.LoginData) error {
	// 检查用户名是否存在
	existingUser, err := u.repo.FindByName(user.Name)
	if err != nil && !errors.Is(err, dao.ErrUserNotFound) {
		// 处理查询中的其他错误
		return err
	}

	if existingUser != nil {
		return errors.New("username already exists")
	}

	// 加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashPassword)

	existingUser = &model.User{
		Name:     user.Name,
		Password: user.Password,
	}

	return u.repo.CreateUser(existingUser)
}

func (u *userService) LoginUser(req *model.LoginData) (interface{}, string, error) {
	user, err := u.repo.FindByName(req.Name)
	if err != nil {
		fmt.Println(err.Error())
		return user, "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return user, "", errors.New("invalid password")
	}

	token, err := util.GenerateToken(user.ID, user.Name, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, "", err
	}

	return user, token, nil
}
