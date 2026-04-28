package models

import (
	"gorm.io/gorm"
)

// User はユーザーの認証情報と基本属性を管理する構造体である
type User struct {
	gorm.Model
	// ログイン時に識別子として使用する一意のユーザー名（またはメールアドレス）
	Username string `gorm:"unique;not null" json:"username"`
	
	// パスワードはセキュリティのため、JSONとしてフロントエンドに返却されないよう制御する
	Password string `gorm:"not null" json:"-"`
	
	// 必要に応じて追加する属性（表示名など）
	DisplayName string `json:"display_name"`
}