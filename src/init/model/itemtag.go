package model

import (
	"gorm.io/gorm"
)

// 商品タグ
type ItemTag struct {
	gorm.Model
	TagName string
	RecommendItem []RecommendItem `gorm:"foreignKey:ItemTagID"`
}