package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 或其他数据库驱动（如 postgres, sqlite 等）
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// sqlx连接
// var DB *sqlx.DB
//
//	func init() {
//		db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/mywork?charset=utf8mb4&parseTime=True&loc=Local")
//		if err != nil {
//			log.Fatalf("数据库连接失败: %v", err)
//		}
//
//		DB = db
//	}
//
// gorm连接
var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/mywork?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println("db 连接失败", err)
	}
	DB = db
}
