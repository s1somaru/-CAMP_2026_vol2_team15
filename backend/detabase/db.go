package database

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 本来は .env から読み込むのが理想だが、まずは直接書いてテストしてもよい
	// ds := "postgres://user:password@hostname:5432/dbname"
	dsn := os.Getenv("DATABASE_URL") 

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("共通DBへの接続に失敗した")
	}

	fmt.Println("RenderのPostgreSQLに接続成功した")
}