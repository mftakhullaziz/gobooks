package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mftakhullaziz/gorest/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file")
	}

	// Config
	UserDb := os.Getenv("DB_USER")
	PassDb := os.Getenv("DB_PASSWORD")
	HostDb := os.Getenv("DB_HOST")
	NameDb := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", UserDb, PassDb, HostDb, NameDb)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create connection to database")
	}

	// Isi model disini untuk auto migrate ke SQL
	err = db.AutoMigrate(&entity.Book{}, &entity.User{})
	if err != nil {
		return nil
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	SqlDb, err := db.DB()

	if err != nil {
		panic("Failed to close connection")
	}

	err = SqlDb.Close()
	if err != nil {
		return
	}
}
