// Package dao internal/dao/user.go
package dao

import (
	"go1/internal/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d *UserDAO) Create(user *model.User) error {
	return d.db.Create(user).Error
}

func (d *UserDAO) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetByID 根据ID获取用户
func (d *UserDAO) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := d.db.First(&user, id).Error
	return &user, err
}

// Update 更新用户信息
func (d *UserDAO) Update(user *model.User) error {
	return d.db.Model(user).Updates(map[string]interface{}{
		"nickname": user.Nickname,
		"email":    user.Email,
	}).Error
}

// GetAll 获取所有用户信息
func (d *UserDAO) GetAll() ([]model.User, error) {
	var users []model.User
	err := d.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
