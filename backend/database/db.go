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
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".envが見つからないため環境変数を直接参照する")
	}

	// 環境変数 DATABASE_URL を取得
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URLが設定されていない")
	}

	var connErr error
	DB, connErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connErr != nil {
		panic("共通DBへの接続に失敗した: " + connErr.Error())
	}

	fmt.Println("RenderのPostgreSQLに接続成功した")
}