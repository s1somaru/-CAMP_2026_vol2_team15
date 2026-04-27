package handlers

import (
	"go-server/models" // プロジェクトのモジュール名（go-server/models）に合わせてください
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCompany は新しい企業情報を登録するAPIです
func CreateCompany(c *gin.Context) {
	var req models.Company

	// リクエストのJSONをCompany構造体にバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不正なリクエストデータです: " + err.Error()})
		return
	}

	// TODO: ここで req のデータを repository を通じてデータベース（DB）に保存する処理を書きます

	// 一旦、受け取ったデータをそのままレスポンスとして返して成功を確認できるようにします
	c.JSON(http.StatusCreated, gin.H{
		"message": "企業情報の登録に成功しました",
		"data":    req,
	})
}

// GetCompanies は登録されているすべての企業情報を取得するAPIです
func GetCompanies(c *gin.Context) {
	// TODO: repository を通じてデータベース（DB）から企業一覧を取得する処理を書きます

	// 開発用の仮データ（テストやフロントエンド作業用）
	mockData := []models.Company{
		{
			ID:                "1",
			CompanyName:       "株式会社A（IT・通信）",
			Status:            "エントリー済み",
			LatestMemo:        "次回：4/25 14:00 オンライン説明会",
			ProgressStep:      "一次",
			ShortMemo:         "社員さんの雰囲気がとても温かかった。自分に合ってそう！",
			EntryDate:         "2026/04/01",
			ESDeadline:        "2026/04/20",
			InterviewMemo:     "・ガクチカの深掘り\n・なぜこの業界なのか\n・逆質問の内容（3つ）",
			EntrySite:         "マイナビ",
			RequiredDocuments: "履歴書、成績証明書",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "企業一覧の取得に成功しました",
		"data":    mockData,
	})
}

// GetCompaniesByID は使わず、一覧すべてを取得するAPIとして変更
func GetAllCompanies(c *gin.Context) {
	// TODO: repository を通じてデータベース（DB）から企業一覧を全て取得する処理を書きます

	// 開発用の仮データ（複数件の配列にして全て返すよう変更）
	mockData := []models.Company{
		{
			ID:          "1",
			CompanyName: "株式会社 サンプル1",
			Status:      "選考中",
		},
		{
			ID:          "2",
			CompanyName: "株式会社 サンプル2",
			Status:      "エントリー済み",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "企業一覧すべての取得に成功しました",
		"data":    mockData,
	})
}
