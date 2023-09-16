package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	_db *gorm.DB
)

func ConnectDb(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //命名策略，不要加s
		},
	})
	if err != nil {
		fmt.Println("attention!! 没有连接上数据库")
		panic(err)
	}
	_db = db

	//创建数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20) //设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //设置数据库打开连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Second * 30) //设置连接可以重用的最长时间
	
	//进行MySQL主从的配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
			// `db2` 作为 sources，`db3`、`db4` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                         // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均衡策略
		}))
	migrate()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}