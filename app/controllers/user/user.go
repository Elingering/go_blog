package user

import (
	"bolg/app/models"
	"bolg/app/service"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/jinzhu/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	//自动检查 User 结构是否变化，变化则进行迁移
	service.DB.AutoMigrate(&models.User{})
    service.DB.Create(&models.User{Name: name, Password:password})
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"password": password,
	})
}