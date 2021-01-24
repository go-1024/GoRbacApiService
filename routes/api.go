package routes

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/controllers/admin"
	admin2 "gorbac/app/middlewares/admin"
)

// 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组
	// 后台 模块
	Admin := router.Group("/admin")
	{
		Admin.Use(admin2.Admin())
		// 登录系统
		Admin.POST("/user/login", admin.LoginHandler)
		// 退出系统
		Admin.POST("/user/logout", admin.LogoutHandler)
		// 获取用户信息
		Admin.POST("/user/info", admin.UserInfoHandler)
		// 首页统计
		Admin.POST("/dashboard/statistics", admin.StatisticsHandler)
		// 获取登录日志
		Admin.POST("/dashboard/login_log", admin.LoginLogHandler)
		// 获取操作日志
		Admin.POST("/dashboard/operation_log", admin.OperationLogHandler)
		// 修改登录密码
		Admin.POST("/dashboard/reset_password", admin.ResetPasswordHandler)
	}
}
