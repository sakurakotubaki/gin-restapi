package main

import (
	"main/internal/database"
	"main/models"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースの初期化
	database.InitDB()

	// Auto Migrate the schema
	err := database.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		panic("Failed to migrate database")
	}

	// Ginサーバーの初期化
	server := gin.Default()

	// ルートの登録
	routes.RegisterRoutes(server)

	// サーバーの起動
	server.Run(":8080")
}
