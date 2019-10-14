package topic

import (
	"bolg/app/models"
	"bolg/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	//自动检查 Topic 结构是否变化，变化则进行迁移
	service.DB.AutoMigrate(&models.Topic{})
}

//新增话题
func Store(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	categoryId, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	userId := int64(1)
	res := service.DB.Create(&models.Topic{Title: title, Body: body, CategoryId: categoryId, UserId: userId})
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

//获取话题列表
func Index(c *gin.Context) {
	var topic []models.Topic
	res := service.DB.Find(&topic)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
		"data":   topic,
	})
}

//话题详情
func Detail(c *gin.Context) {
	id := c.Param("id")
	var topic models.Topic
	res := service.DB.Where("id = ? ", id).First(&topic)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   "200",
		"data":   topic,
	})
}

//编辑话题
func Edit(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body")
	categoryId, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	var topic models.Topic
	res := service.DB.First(&topic, id)
	checkErr(res.Error)
	topic.Title = title
	topic.Body = body
	topic.CategoryId = categoryId
	res = res.Save(&topic)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

//删除话题
func Delete(c *gin.Context) {
	id := c.Param("id")
	var topic models.Topic
	res := service.DB.Delete(&topic, "id = ?", id)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "delete topic",
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
