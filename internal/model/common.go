// internal/model/common.go
package model

// 分页请求参数
type PageRequest struct {
	Page     int `form:"page" json:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required,min=1,max=100"`
}

// 分页响应数据
type PageResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
