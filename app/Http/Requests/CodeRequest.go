package Requests

import (
	"gopkg.in/go-playground/validator.v8"
)

type CodeRequest struct {
	Phone   string `form:"phone" binding:"required,phone"`
	CodeKey string `form:"code_key" binding:"required"`
	Code    string `form:"code" binding:"required,len=5"`
}

func (r *CodeRequest) GetError(err validator.ValidationErrors) string {
	// 这里的 "CodeRequest.Phone" 索引对应的是模型的名称和字段
	if val, exist := err["CodeRequest.Phone"]; exist {
		if val.Field == "Phone" {
			switch val.Tag {
			case "required":
				return "请输入Phone"
			case "phone":
				return "Phone格式不正确"
			}
		}
	}
	if val, exist := err["CodeRequest.CodeKey"]; exist {
		if val.Field == "CodeKey" {
			switch val.Tag {
			case "required":
				return "请输入CodeKey"
			}
		}
	}
	if val, exist := err["CodeRequest.Code"]; exist {
		if val.Field == "Code" {
			switch val.Tag {
			case "required":
				return "请输入Code"
			case "len":
				return "Code长度为6个字符"
			}
		}
	}
	return "参数错误"
}
