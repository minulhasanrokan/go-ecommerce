package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, database)

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		return nil, err
	}

	log.Default().Println("Successfully connected to mySql database")

	return db, nil
}
