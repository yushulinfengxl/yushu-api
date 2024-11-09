package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connected(dsn string, tablePrefix string) (*gorm.DB, error) {
	// 配置 GORM 数据库连接池
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 可配置日志等级以便调试，选择 GORM 日志级别
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 表名前缀，`User` 对应的表名是 `t_users`
			SingularTable: false,       // 使用单数表名，启用该选项后，`User` 表将是`user`
			NoLowerCase:   false,       // 禁用小写转换
		},
	})
}
