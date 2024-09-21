package model

import (
	"gorm.io/gorm"
)

// 紹介商品リスト
type RecommendItem struct {
	gorm.Model
	Product_name string `gorm:"size:100"`
	Detail string
	Category string
	Link string
	Image string
	Price int
	ItemTagID uint
	Tweets *[]Tweets `gorm:"foreignKey:RecommendItemID"`
}