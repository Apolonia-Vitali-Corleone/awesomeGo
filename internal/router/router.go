// Package router internal/router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"go1/internal/middleware"
	"go1/internal/model"
	"go1/internal/service"
	"go1/pkg/e"
	"go1/pkg/response"
)

// RegisterRequest 用户注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginRequest 用户登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InitRouter(userService *service.UserService) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 公开路由
	public := r.Group("/api/v1")
	{
		// 注册处理器
		public.POST("/register", func(c *gin.Context) {
			var req RegisterRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				response.Error(c, e.InvalidParams, err.Error())
				return
			}

			err := userService.Register(&model.User{
				Username: req.Username,
				Password: req.Password,
				Nickname: req.Nickname,
				Email:    req.Email,
			})

			if err != nil {
				response.Error(c, e.ERROR, err.Error())
				return
			}

			response.Success(c, nil)
		})

		// 登录处理器
		public.POST("/login", func(c *gin.Context) {
			var req LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				response.Error(c, e.InvalidParams, err.Error())
				return
			}

			token, err := userService.Login(req.Username, req.Password)
			if err != nil {
				response.Error(c, e.ERROR, err.Error())
				return
			}

			response.Success(c, gin.H{
				"token": token,
			})
		})
	}

	// 需要认证的路由
	auth := r.Group("/api/v1")
	auth.Use(middleware.JWT())
	{
		// 获取用户信息
		auth.GET("/user/info", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			userInfo, err := userService.GetUserInfo(userID.(uint))
			if err != nil {
				response.Error(c, e.ERROR, err.Error())
				return
			}
			response.Success(c, userInfo)
		})

		// 更新用户信息
		auth.PUT("/user/info", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			var userInfo model.User
			if err := c.ShouldBindJSON(&userInfo); err != nil {
				response.Error(c, e.InvalidParams, err.Error())
				return
			}
			userInfo.ID = userID.(uint)

			err := userService.UpdateUserInfo(&userInfo)
			if err != nil {
				response.Error(c, e.ERROR, err.Error())
				return
			}
			response.Success(c, nil)
		})
	}

	return r
}
