package routes

import (
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

	return r
}
