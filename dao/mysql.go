package dao

import (
	"gorbac/app/models"
	"gorbac/app/models/account"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 建立数据库连接池
	db := models.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(100)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(5*60) * time.Second)
	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	// 自动迁移
	db.AutoMigrate(
		&account.Model{},
	)
}
