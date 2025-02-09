// Package response pkg/response/response.go
package response

import (
	"github.com/gin-gonic/gin"
	"go1/pkg/e"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

// PageResult 分页响应
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// Page 分页响应
func Page(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	c.JSON(200, Response{
		Code:    e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data: PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	})
}