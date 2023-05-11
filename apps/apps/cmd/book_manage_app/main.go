package main

import (
    "context"
    "log"
    "book_manage_app/ent"

    "entgo.io/ent/dialect"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // インメモリーのSQLiteデータベースを持つent.Clientを作成します。
    client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    // 出力します。
	task1, err := client.User.Create().Save(ctx)
    if err != nil {
        log.Fatalf("failed creating a todo: %v", err)
    }
    log.Println(task1)
}



/*
import (
	_ "book_manage_app/docs"
	"book_manage_app/internal/api/routes"
	"book_manage_app/internal/db"
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
*/