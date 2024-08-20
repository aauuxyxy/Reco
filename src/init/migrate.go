package migrate

import (
	"Reco-server/src/init/model"

	"gorm.io/gorm"
)

func MySQL_Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Users{}, &model.Tweets{}, &model.RecommendItem{}, &model.Retweets{}, &model.Follow{})
}