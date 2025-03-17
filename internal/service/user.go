// Package service internal/service/user.go
package service

import (
	"errors"
	"go1/internal/dao"
	"go1/internal/model"
	"go1/pkg/util"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

func (s *UserService) Register(user *model.User) error {
	// 密码加密
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userDAO.Create(user)
}

// Login 登录
func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.userDAO.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// 生成JWT token
	token, err := util.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID uint) (*model.User, error) {
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUserInfo 更新用户信息
func (s *UserService) UpdateUserInfo(user *model.User) error {
	if user.ID == 0 {
		return errors.New("invalid user id")
	}
	return s.userDAO.Update(user)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.userDAO.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
