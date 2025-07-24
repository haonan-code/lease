// Package recover_middleware 提供全局异常恢复中间件
// 创建者：Done-0
// 创建时间：2025-05-10
package recover_middleware

import (
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"jank.com/jank_blog/internal/global"
)

// InitRecover 初始化全局异常恢复中间件
// 返回值：
//   - echo.MiddlewareFunc: Echo 框架中间件函数
func InitRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4096, // 堆栈大小
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			stackSize := 4096
			var buf []byte
			for {
				buf = make([]byte, stackSize)
				n := runtime.Stack(buf, false)
				if n < stackSize {
					buf = buf[:n]
					break
				}
				stackSize *= 2
			}

			// 将完整的堆栈轨迹信息记录到日志
			global.SysLog.WithFields(map[string]interface{}{
				"stack_trace": string(buf),
			}).Errorf("发生运行时异常: %v", err)
			return nil
		},
	})
}
