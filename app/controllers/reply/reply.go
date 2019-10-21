package reply

import (
	"bolg/app/models"
	"bolg/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	//自动检查 Reply 结构是否变化，变化则进行迁移
	service.DB.AutoMigrate(&models.Reply{})
}

//新增回复
func Store(c *gin.Context) {
	content := c.PostForm("content")
	topicId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	userId := int64(1)
	res := service.DB.Create(&models.Reply{Content: content, TopicId: topicId, UserId: userId})
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "",
	})
}

//获取文章的回复列表
func TopicIndex(c *gin.Context) {
	//id := c.Param("id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var topic models.Topic
	service.DB.Where("id = ?", id).First(&topic)
	res := service.DB.Debug().Model(&topic).Related(&topic.Reply)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
		"data":   topic.Reply,
	})
}

//获取用户的回复列表
func UserIndex(c *gin.Context) {
	//id := c.Param("id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var user models.User
	service.DB.Where("id = ?", id).First(&user)
	res := service.DB.Debug().Model(&user).Related(&user.Reply)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
		"data":   user.Reply,
	})
}

//删除回复
func Delete(c *gin.Context) {
	//id := c.Param("id")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var reply models.Reply
	res := service.DB.Delete(&reply, "id = ?", id)
	checkErr(res.Error)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "delete reply",
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
