package db

import (
	"book_manage_app/ent"
	"book_manage_app/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
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
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		MaxIdleConns: serverConfig.DataBase.MaxIdleConns,
		MaxOpenConns: serverConfig.DataBase.MaxOpenConns,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v\n", err)
	}
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}
