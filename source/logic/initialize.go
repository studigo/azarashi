package logic

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Model の初期化を行う.
func Initialize() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	instance, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return fmt.Errorf("initialize failed: %s", err)
	}

	db = instance

	return nil
}
