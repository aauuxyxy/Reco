package main

import (
	migrate "Reco-server/src/init"
	"Reco-server/src/init/model"
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
    r.GET("/all_users", func(c *gin.Context) {
        var users []model.Users
        db.Find(&users)

        c.JSON(http.StatusOK, users)
    })

    r.Run()
}