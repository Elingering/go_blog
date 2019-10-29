package Models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Reply    []Reply
	Name     string
	Password string
	Phone    string
	Age      int8
	Email    string
}
