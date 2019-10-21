package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Reply    []Reply
	Name     string
	Password string
	Age      int8
	Email    string
}
