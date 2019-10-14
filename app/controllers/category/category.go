package category

import (
	"bolg/app/models"
	"bolg/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	//自动检查 Category 结构是否变化，变化则进行迁移
	service.DB.AutoMigrate(&models.Category{})
}

//新增分类
func Store(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	res := service.DB.Create(&models.Category{Name: name, Description: description})
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

//获取分类列表
func Index(c *gin.Context) {
	var category []models.Category
	res := service.DB.Find(&category)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
		"data":   category,
	})
}

//编辑分类
func Edit(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	description := c.PostForm("description")
	var category models.Category
	res := service.DB.First(&category, id)
	checkErr(res.Error)
	category.Name = name
	category.Description = description
	res = res.Save(&category)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

//删除分类
func Delete(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	res := service.DB.Delete(&category, "id = ?", id)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
