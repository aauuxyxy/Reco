package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Auth_id string `gorm:"not null"`
	Name string `gorm:"size:100;not null"`
	Bio string `gorm:"size:280;default:null"`
	Birth_date time.Time `gorm:"default:null"`
	Profile_image string `gorm:"default:null"`
	Tweets []Tweets `gorm:"foreignKey:UsersID;default:null"`
	Retweets []Retweets `gorm:"foreignKey:UsersID;default:null"`
	Follow []Follow `gorm:"foreignKey:UsersID;default:null"`
}