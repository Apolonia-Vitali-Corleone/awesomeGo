// internal/model/operation_log.go
package model

import (
	"gorm.io/gorm"
)

type OperationLog struct {
	gorm.Model
	UserID    uint   `gorm:"index" json:"user_id"`
	IP        string `gorm:"size:64" json:"ip"`
	Method    string `gorm:"size:20" json:"method"`
	Path      string `gorm:"size:255" json:"path"`
	Status    int    `json:"status"`
	Latency   int64  `json:"latency"` // 延迟时间(ms)
	UserAgent string `gorm:"size:255" json:"user_agent"`
	Request   string `gorm:"type:text" json:"request"`
	Response  string `gorm:"type:text" json:"response"`
}
