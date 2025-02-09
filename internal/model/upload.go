// internal/model/upload.go
package model

import (
	"gorm.io/gorm"
)

type Upload struct {
	gorm.Model
	UserID   uint   `gorm:"index" json:"user_id"`
	FileName string `gorm:"size:255" json:"file_name"`
	FilePath string `gorm:"size:255" json:"file_path"`
	FileSize int64  `json:"file_size"`
	FileType string `gorm:"size:64" json:"file_type"`
	Status   int    `gorm:"type:tinyint(1);default:1" json:"status"` // 0-禁用 1-启用
}
