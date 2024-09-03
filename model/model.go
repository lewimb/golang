package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	Username   string `json:"username" gorm:"uniqueIndex;not null;unique"`
	Password   string `json:"password" gorm:"not null"`
	Fullname   string `json:"fullname"`
	Desc       string `json:"desc"`
	ProfilePic string `json:"profilePic"`
	Photos     []Photo `gorm:"foreignKey:UserID;references:ID"`
}

type Photo struct {
	gorm.Model
	UserID    uint `json:"userId" gorm:"not null"`
	URL       string `json:"url"`
	Caption   string `json:"caption"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
