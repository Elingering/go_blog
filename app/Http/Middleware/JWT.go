package Middleware

import (
	"bolg/app/Services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if "" == tokenString {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"code":    401,
				"message": "Have no token",
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		//redis 黑名单
		val, err := Services.TX.Exists(tokenString).Result()
		println(val)
		if err != nil {
			panic(err)
		}
		if 1 == val {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"code":    401,
				"message": "Token timeout",
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		type MyCustomClaims struct {
			Foo string `json:"foo"`
			jwt.StandardClaims
		}
		func() {
			token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("AllYourBase"), nil
			})
			if token == nil {
				// 验证不通过，不再调用后续的函数处理
				c.Abort()
				msg := fmt.Sprint("Couldn't handle this token:", err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "error",
					"code":    401,
					"message": msg,
				})
				// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
				return
			}
			if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
				// 验证通过，会继续访问下一个中间件
				c.Next()
			} else {
				// 验证不通过，不再调用后续的函数处理
				c.Abort()
				msg := fmt.Sprint("Couldn't handle this token:", err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "error",
					"code":    401,
					"message": msg,
				})
				// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
				return
			}
		}()
	}
}
