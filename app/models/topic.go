package models

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	Title      string
	Body       string `gorm:"type:text"`
	UserId     int64
	CategoryId int64
}
