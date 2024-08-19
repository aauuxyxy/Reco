package migrate

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ユーザーテーブル
type Users struct {
	gorm.Model
	Auth_id string
	Name string `gorm:"size:100"`
	Bio string `gorm:"size:280"`
	Birth_date time.Time
	Profile_image string
	Tweets []Tweets `gorm:"foreignKey:UsersID"`
	Retweets []Retweets `gorm:"foreignKey:UsersID"`
	Follow []Follow `gorm:"foreignKey:UsersID"`
}

// ツイート
type Tweets struct {
	gorm.Model
	content string `gorm:"size:280"`
	RecommendItem RecommendItem `gorm:"foreignKey:RecommendItemID"`
	Image string
	reply_to uint //リプライの場合に返信先IDを格納
	quote_tweet_id uint //引用ツイートの場合に引用元IDを格納
	Retweets []Retweets `gorm:"foreignKey:TweetsID"`
	RecommendItemID uint
	UsersID uint
}

// 紹介商品リスト
type RecommendItem struct {
	gorm.Model
	Product_name string `gorm:"size:100"`
	Detail string
	Category string
	Link string
	Image string
	Tweets []Tweets `gorm:"foreignKey:RecommendItemID"`
}

// リツイート関連テーブル
type Retweets struct {
	UsersID uint
	TweetsID uint
}

// フォロー
type Follow struct {
	UsersID uint
	Follower Users
}

func MySQL_Migrate() {
	dsn := "root@tcp(mysql:3306)/first?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Users{}, &Tweets{}, &RecommendItem{}, &Retweets{}, &Follow{})

	fmt.Println("Migration Done")
}