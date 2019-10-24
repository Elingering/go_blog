package routes

import (
	"bolg/app/Http/Controllers/category"
	"bolg/app/Http/Controllers/reply"
	"bolg/app/Http/Controllers/topic"
	"bolg/app/Http/Controllers/user"
	"bolg/app/Http/Middleware"
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

	auth := r.Group("")
	auth.Use(Middleware.Jwt())
	{
		//用户退出
		auth.POST("/sign-out", user.Logout)
		//用户信息
		auth.GET("/user", user.Index)
		//编辑用户
		auth.PATCH("/user", user.Edit)

		//创建分类
		auth.POST("/category", category.Store)
		//分类列表
		auth.GET("/category", category.Index)
		//编辑分类
		auth.PATCH("/category/:id", category.Edit)
		//删除分类
		auth.DELETE("/category/:id", category.Delete)

		//创建话题
		auth.POST("/topic", topic.Store)
		//话题列表
		auth.GET("/topic", topic.Index)
		//话题详情
		auth.GET("/topic/:id", topic.Detail)
		//编辑话题
		auth.PATCH("/topic/:id", topic.Edit)
		//删除话题
		auth.DELETE("/topic/:id", topic.Delete)

		//新增回复
		auth.POST("/topic/:id/reply", reply.Store)
		//获取文章的回复列表
		auth.GET("/topic/:id/reply", reply.TopicIndex)
		//获取用户的回复列表
		auth.GET("/user/:id/reply", reply.UserIndex)
		//删除回复
		auth.DELETE("/reply/:id", reply.Delete)
	}

	return r
}
