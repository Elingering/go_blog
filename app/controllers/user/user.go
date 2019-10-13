package user

import (
	"bolg/app/models"
	"bolg/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	//自动检查 User 结构是否变化，变化则进行迁移
	service.DB.AutoMigrate(&models.User{})
}

//注册接口
func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	//string转int64
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 8)
	email := c.PostForm("email")
	res := service.DB.Create(&models.User{Name: name, Password: password, Age: int8(age), Email: email})
	if res.Error != nil {
		panic(res.Error)
	}
	c.JSON(http.StatusOK, gin.H{
		"name":     name,
		"password": password,
	})
}

//登录接口
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var user models.User
	res := service.DB.Where("email = ? and password = ?", email, password).Find(&user)
	checkErr(res.Error)
	res.Scan(&user)
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

//退出接口
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "logout",
	})
}

//获取用户信息
func Index(c *gin.Context) {
	id := c.Query("id")
	var user models.User
	res := service.DB.First(&user, id)
	checkErr(res.Error)
	res.Scan(&user)
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"age":   user.Age,
		"email": user.Email,
	})
}

//编辑用户信息
func Edit(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 8)
	email := c.PostForm("email")
	var user models.User
	res := service.DB.First(&user, id)
	checkErr(res.Error)
	user.Name = name
	user.Age = int8(age)
	user.Email = email
	res.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"age":   user.Age,
		"email": user.Email,
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
