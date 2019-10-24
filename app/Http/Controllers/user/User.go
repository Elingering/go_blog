package user

import (
	"bolg/app/Http/Requests"
	"bolg/app/Models"
	"bolg/app/Services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	//自动检查 User 结构是否变化，变化则进行迁移
	Services.DB.AutoMigrate(&Models.User{})
}

//注册接口
func Register(c *gin.Context) {
	var UserRequest Requests.UserRequest
	if err := c.ShouldBind(&UserRequest); err == nil {
		name := c.PostForm("name")
		password := c.PostForm("password")
		//string转int64
		age, _ := strconv.ParseInt(c.PostForm("age"), 10, 8)
		email := c.PostForm("email")
		res := Services.DB.Create(&Models.User{Name: name, Password: password, Age: int8(age), Email: email})
		checkErr(res.Error)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"code":    200,
			"message": "",
		})
	} else {
		// 验证错误
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"code":    400,
			"message": UserRequest.GetError(err.(validator.ValidationErrors)), // 注意这里要将 err 进行转换
			//"message": err.Error(), // 注意这里要将 err 进行转换
		})
	}
}

//登录接口
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	var user Models.User
	res := Services.DB.Where("email = ? and password = ?", email, password).First(&user)
	checkErr(res.Error)
	//生成token
	mySigningKey := []byte("AllYourBase")
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	// Create the Claims
	expired := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: expired,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	checkErr(err)
	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"code":       200,
		"token":      ss,
		"expires_at": expired,
	})
}

//退出接口
func Logout(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	var nx int64
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		nx = claims.ExpiresAt
	} else {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"code":    500,
			"message": "logout error",
		})
	}
	//redis 黑名单
	Services.TX.SetNX(tokenString, 1, time.Duration(nx-time.Now().Unix())*time.Second)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    200,
		"message": "logout",
	})
}

//获取用户信息
func Index(c *gin.Context) {
	id := c.Query("id")
	var user Models.User
	res := Services.DB.First(&user, id)
	checkErr(res.Error)
	res.Scan(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   200,
		"data":   user,
	})
}

//编辑用户信息
func Edit(c *gin.Context) {
	id := 1
	name := c.PostForm("name")
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 8)
	email := c.PostForm("email")
	var user Models.User
	res := Services.DB.First(&user, id)
	checkErr(res.Error)
	user.Name = name
	user.Age = int8(age)
	user.Email = email
	res = res.Save(&user)
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
