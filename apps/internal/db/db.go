package db

import (
	"book_manage_app/ent"
	"book_manage_app/internal/config"
	"fmt"
	"log"
	"os"
	"time"

	"entgo.io/ent/dialect/sql"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DBName       string
	MaxIdleConns int
	MaxOpenConns int
}

func NewDB() (*ent.Client, error) {
	serverConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dbConfig := DatabaseConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		MaxIdleConns: serverConfig.DataBase.MaxIdleConns,
		MaxOpenConns: serverConfig.DataBase.MaxOpenConns,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName)

	drv, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("データベースへの接続に失敗しました: %v\n", err)
    }

    db := drv.DB()
    db.SetMaxIdleConns(dbConfig.MaxIdleConns)
    db.SetMaxOpenConns(dbConfig.MaxOpenConns)
    db.SetConnMaxLifetime(time.Hour)
    return ent.NewClient(ent.Driver(drv)), nil
}
