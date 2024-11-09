package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
	"yushu/box/config"
	"yushu/box/utility/singleton"
)

var dbLazySingleton = &singleton.Lazy{}

func NewDB() *gorm.DB {
	return (*dbLazySingleton.Instance(gorm.DB{})).(*gorm.DB)
}

func New() *gorm.DB {
	//conn := &gorm.DB{}
	defer func() {
		r := recover()
		if r != nil {
			log.Println("数据库连接失败recover", r)
		}
	}()

	// 获取配置
	mysqlConf := config.New().Database.Mysql

	db := NewDB()
	dsn := ""
	isValid := false
	// 连接字符串
	for _, item := range mysqlConf.List {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
			item.User, item.Password, item.Host, item.Port, item.Database, item.Timeout)
		db2, err := Connected(dsn, mysqlConf.TablePrefix)
		if err == nil {
			isValid = true
		}
		db = db2
	}

	/*
	 */
	// 设置 MySQL 数据库连接信息
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// 配置 GORM 数据库连接池
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	// 可配置日志等级以便调试，选择 GORM 日志级别
	//	Logger: logger.Default.LogMode(logger.Info),
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   config.New().TablePrefix, // 表名前缀，`User` 对应的表名是 `t_users`
	//		SingularTable: false,                    // 使用单数表名，启用该选项后，`User` 表将是`user`
	//		NoLowerCase:   false,                    // 禁用小写转换
	//	},
	//})
	//if err != nil {
	//	fmt.Println("Failed to connect to database:", err)
	//	return
	//}

	// 获取通用数据库接口以配置连接池
	if !isValid {
		panic("数据库连接失败")
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get generic database interface:", err)

	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxIdleTime(time.Hour) // 空闲连接最大存活时间
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大存活时间
	/*

	 */
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
	//	conf.List[0].User, conf.List[0].Password, conf.List[0].Host, conf.Port, configs.Database, configs.Timeout)
	return nil
}
