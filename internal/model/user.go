// Package model internal/model/
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
	Role     string `gorm:"size:64" json:"role"`
}