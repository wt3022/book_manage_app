package main

import (
	_ "book_manage_app/docs"
	"book_manage_app/internal/db"
	"book_manage_app/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title		example
// @version	1.0
// @BasePath	/
// @in			header
// @name		hoge
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("サーバー起動時に.envファイルが読み込めませんでした。")
	}

	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	routes.InitRoutes(e, db)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
}
