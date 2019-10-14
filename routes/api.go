package routes

import (
	"bolg/app/controllers/category"
	"bolg/app/controllers/topic"
	"bolg/app/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutes() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//用户注册
	r.POST("/sign-up", user.Register)
	//用户登录
	r.POST("/sign-in", user.Login)
	//用户退出
	r.POST("/sign-out", user.Logout)
	//用户信息
	r.GET("/user", user.Index)
	//编辑用户
	r.PATCH("/user", user.Edit)

	//创建分类
	r.POST("/category", category.Store)
	//分类列表
	r.GET("/category", category.Index)
	//编辑分类
	r.PATCH("/category/:id", category.Edit)
	//删除分类
	r.DELETE("/category/:id", category.Delete)

	//创建话题
	r.POST("/topic", topic.Store)
	//话题列表
	r.GET("/topic", topic.Index)
	//话题详情
	r.GET("/topic/:id", topic.Detail)
	//编辑话题
	r.PATCH("/topic/:id", topic.Edit)
	//删除话题
	r.DELETE("/topic/:id", topic.Delete)

	return r
}
