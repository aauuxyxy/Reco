package model

import (
	"gorm.io/gorm"
)

// ツイート
type Tweets struct {
	gorm.Model
	content string `gorm:"size:280"`
	RecommendItem RecommendItem `gorm:"foreignKey:RecommendItemID"`
	Image *string
	reply_to *uint //リプライの場合に返信先IDを格納
	quote_tweet_id *uint //引用ツイートの場合に引用元IDを格納
	Retweets *[]Retweets `gorm:"foreignKey:TweetsID"`
	RecommendItemID uint
	UsersID uint
}