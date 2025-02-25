// main.go
package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go1/config"
	"go1/internal/dao"
	"go1/internal/router"
	"go1/internal/service"
	"go1/pkg/database"
	"log"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		fmt.Println("任务被取消")
	}
}
func main() {

	// 加载配置
	var cfg config.Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
	}

	// 初始化MySQL连接
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// 初始化Redis连接
	rdb, err := database.InitRedis(cfg.Redis)
	if err != nil {
		log.Fatalf("Error connecting to redis: %s", err)
	}
	_ = rdb // 暂时不使用redis，避免未使用变量的警告

	// 初始化DAO层
	userDAO := dao.NewUserDAO(db)

	// 初始化Service层
	userService := service.NewUserService(userDAO)

	// 初始化路由
	r := router.InitRouter(userService)

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
