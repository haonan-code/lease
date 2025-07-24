// Package middleware 提供中间件集成和初始化功能
package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"lease/configs"
	"lease/internal/global"
	cors_middleware "lease/internal/middleware/cors"
	error_middleware "lease/internal/middleware/error"
	logger_middleware "lease/internal/middleware/logger"
	recover_middleware "lease/internal/middleware/recover"
	secure_middleware "lease/internal/middleware/secure"
	swagger_middleware "lease/internal/middleware/swagger"
)

// New 初始化并注册所有中间件
// 参数：
//   - app: gin 实例
func New(app *gin.Engine) {
	// 设置全局错误处理
	app.Use(error_middleware.InitError())
	// 配置 CORS 中间件
	app.Use(cors_middleware.InitCORS())
	// 全局请求 ID 中间件
	//app.Use(middleware.RequestID())
	app.Use(requestid.New())
	// 日志中间件
	app.Use(logger_middleware.InitLogger())
	// 配置 xss 防御中间件
	app.Use(secure_middleware.InitXss())
	// 配置 csrf 防御中间件
	app.Use(secure_middleware.InitCSRF())
	// 全局异常恢复中间件
	app.Use(recover_middleware.InitRecover())

	// Swagger中间件初始化
	initSwagger(app)
}

// initSwagger 根据配置初始化 Swagger 文档中间件
// 参数：
//   - app: gin实例
func initSwagger(app *gin.Engine) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		global.SysLog.Errorf("加载 Swagger 配置失败: %v", err)
		return
	}

	switch cfg.SwaggerConfig.SwaggerEnabled {
	case "true":
		app.Use(swagger_middleware.InitSwagger())
		global.SysLog.Info("Swagger 已启用")
	default:
		global.SysLog.Info("Swagger 已禁用")
	}
}
