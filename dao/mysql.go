package dao

import (
	_ "github.com/go-sql-driver/mysql" // 或其他数据库驱动（如 postgres, sqlite 等）
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/mywork?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	DB = db
}
