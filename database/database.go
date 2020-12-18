package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go_web/pkg/logger"
	"time"
)

// 数据库
var DB *sql.DB

func Initialize() {
	//初始化数据库
	initDB()
	// 初始化数据表
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "192.168.10.10:3306",
		Net:                  "tcp",
		DBName:               "go_web",
		AllowNativePasswords: true,
	}

	//准备数据库连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = DB.Ping()
	logger.LogError(err)
}

// 创建 articles 数据表
func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
