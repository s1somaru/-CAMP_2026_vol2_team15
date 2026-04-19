package database

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 1. .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".envファイルが見つかりません。環境変数から直接読み込みます。")
	}

	// 2. 環境変数からURLを取得
	dsn := os.Getenv("DATABASE_URL")

	// 3. Postgresに接続
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("共通DBへの接続に失敗した")
	}

	fmt.Println("RenderのPostgreSQLに接続成功した")
}