package Requests

import "gopkg.in/go-playground/validator.v8"

type UserRequest struct {
	Name     string `form:"name" json:"name" binding:"required,len=5"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *UserRequest) GetError(err validator.ValidationErrors) string {
	// 这里的 "LoginRequest.Mobile" 索引对应的是模型的名称和字段
	if val, exist := err["UserRequest.Name"]; exist {
		if val.Field == "Name" {
			switch val.Tag {
			case "required":
				return "请输入Name"
			case "len":
				return "len <> 5"
			}
		}
	}
	if val, exist := err["UserRequest.Password"]; exist {
		if val.Field == "Password" {
			switch val.Tag {
			case "required":
				return "请输入Password"
			}
		}
	}
	return "参数错误"
}
