package Requests

import "gopkg.in/go-playground/validator.v8"

type UserRequest struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required,min=6"`
	Phone    string `form:"phone" binding:"required,phone"`
	Email    string `form:"email" binding:"required,email"`
	Age      int8   `form:"age" binding:"-"`
	Code     string `form:"code" binding:"required,len=6"`
}

func (r *UserRequest) GetError(err validator.ValidationErrors) string {
	// 这里的 "UserRequest.Name" 索引对应的是模型的名称和字段
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
			case "min":
				return "Password至少6位"
			}
		}
	}
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
	if val, exist := err["UserRequest.Email"]; exist {
		if val.Field == "Email" {
			switch val.Tag {
			case "required":
				return "请输入Email"
			case "email":
				return "Email格式不对"
			}
		}
	}
	if val, exist := err["UserRequest.Code"]; exist {
		if val.Field == "Code" {
			switch val.Tag {
			case "required":
				return "请输入Code"
			}
		}
	}
	return "参数错误"
}
