package main

import (
	"gorbac/dao"
	"gorbac/routes"
)

func main() {
	// 初始化 SQL
	dao.SetupDB()
	// 注册路由服务
	routes.RegisterWebRoutes()
}
