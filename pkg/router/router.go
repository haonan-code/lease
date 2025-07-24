// Package router 提供应用程序路由注册功能
package router

import (
	"github.com/gin-gonic/gin"
)

// New @title		Lease API
// @version		1.0
// @description	This is the API documentation for Lease System.
// @host		localhost:9010
// @BasePath	/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description 输入格式: Bearer {token}
// New 函数用于注册应用程序的路由
// 参数：
//   - app: gin 实例
func New(app *gin.Engine) {
	// 创建多版本 API 路由组
	api1 := app.Group("/api/v1")
	//api2 := app.Group("/api/v2")

	// 注册测试相关的路由
	//routers.RegisterTestRoutes(api1, api2)
	// 注册账户相关的路由
	routers.RegisterAccountRoutes(api1)
	// 注册验证相关的路由
	//routers.RegisterVerificationRoutes(api1)
	//// 注册文章相关的路由
	//routers.RegisterPostRoutes(api1)
	//// 注册类目相关的路由
	//routers.RegisterCategoryRoutes(api1)
	//// 注册评论相关的路由
	//routers.RegisterCommentRoutes(api1)
}
