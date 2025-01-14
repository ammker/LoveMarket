package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"lovemarket/setting"
	"time"
)

// Init 初始化MySQL连接
func Init(cfg *setting.MySQLConfig) (err error) {
	// 构造 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("Failed to get database object, got error: %v", err)
		return err
	}

	// 设置连接池配置
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

// Close 关闭MySQL连接
func Close() {
	sqlDB, err := Db.DB() // 获取底层的 sql.DB 对象
	if err != nil {
		log.Fatalf("Failed to get database object, got error: %v", err)
		return
	}
	err = sqlDB.Close() // 关闭连接
	if err != nil {
		log.Fatalf("Failed to close database connection, got error: %v", err)
	}
}
