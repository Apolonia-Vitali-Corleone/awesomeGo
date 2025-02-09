// Package model internal/model/user.go
package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:64" json:"username"`
	Password string `gorm:"size:128" json:"-"`
	Nickname string `gorm:"size:128" json:"nickname"`
	Email    string `gorm:"size:128" json:"email"`
	Role     string `gorm:"size:64;default:user" json:"role"`
	Status   int    `gorm:"type:tinyint(1);default:1" json:"status"` // 0-禁用 1-启用
}
