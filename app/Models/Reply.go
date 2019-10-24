package Models

import "github.com/jinzhu/gorm"

type Reply struct {
	gorm.Model
	Content string `gorm:"type:text"`
	UserId  int64
	TopicId int64
}
