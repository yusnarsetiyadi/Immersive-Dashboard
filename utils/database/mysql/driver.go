package mysql

import (
	"api-alta-dashboard/config"
	class "api-alta-dashboard/features/class/repository"
	logs "api-alta-dashboard/features/log/repository"
	mentee "api-alta-dashboard/features/mentee/repository"
	user "api-alta-dashboard/features/user/repository"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&class.Class{})
	db.AutoMigrate(&mentee.Mentee{})
	db.AutoMigrate(&logs.Log{})
}
