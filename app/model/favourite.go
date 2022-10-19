package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type Favourite struct {
	gorm.Model
	Platform string `gorm:"not null" json:"platform"`
	RoomId   string `gorm:"not null" json:"room_id"`
	Upper    string `json:"upper"`
}

func (Favourite) TableName() string {
	// return
	return "favourite"
}
