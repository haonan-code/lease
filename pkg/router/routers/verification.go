// Package routes 提供路由注册功能
// 创建者：Done-0
// 创建时间：2025-05-10
package routes

import (
	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/pkg/serve/controller/verification"
)

// RegisterVerificationRoutes 注册验证码相关路由
// 参数：
//   - r: Echo 路由组数组，r[0] 为 API v1 版本组
func RegisterVerificationRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/verification")
	accountGroupV1.GET("/sendImgVerificationCode", verification.SendImgVerificationCode)
	accountGroupV1.GET("/sendEmailVerificationCode", verification.SendEmailVerificationCode)
}
