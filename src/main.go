package main

import (
	migrate "Reco-server/src/init"
	"Reco-server/src/init/seeding"
	"Reco-server/src/pkg/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

    // データベースに接続
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // マイグレーション
    migrate.MySQL_Migrate(db)
    seeding.Seeder(db)

    // ルーティング
    r := gin.Default()
    r.GET("/api/data", func(c *gin.Context) {
        response := gin.H{
            "message": "Hello World!",
        }
        c.JSON(http.StatusOK, response)
    })
    r.Run()
}