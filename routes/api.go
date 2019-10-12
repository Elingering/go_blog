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

	return r
}