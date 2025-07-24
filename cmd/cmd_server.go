package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lease/configs"
	"lease/internal/db"
	"lease/internal/logger"
	"lease/internal/middleware"
	"lease/internal/redis"
	"lease/pkg/router"
	"log"
)

func Start() {
	if err := configs.Init(configs.DefaultConfigPath); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
		return
	}

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("获取配置失败: %v", err)
		return
	}

	// 初始化 Logger
	logger.New()

	// 初始化 gin 实例
	app := gin.New()

	// 初始化中间件
	middleware.New(app)

	// 初始化数据库连接并自动迁移模型
	db.New(config)

	// 初始化 Redis 连接
	redis.New(config)

	// 注册路由
	router.New(app)

	// 启动服务
	addr := fmt.Sprintf("%s:%s", config.AppConfig.AppHost, config.AppConfig.AppPort)
	log.Printf("Gin server starting on %s...", addr)
	if err := app.Run(addr); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}

	//app.Logger.Fatal(app.Run(fmt.Sprintf("%s:%s", config.AppConfig.AppHost, config.AppConfig.AppPort)))
}
