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

//获取图像验证码
func GetVCode(c *gin.Context) {
	codeKey, base64stringD := Services.GetVerificationCode()
	c.String(http.StatusOK, codeKey+"\n"+base64stringD)
}

//获取手机验证码
func GetCode(c *gin.Context) {
	var CodeRequest Requests.CodeRequest
	if err := c.ShouldBindQuery(&CodeRequest); err == nil {
		phone := CodeRequest.Phone
		codeKey := CodeRequest.CodeKey
		code := CodeRequest.Code
		res, err := Services.TX.Get(codeKey).Result()
		checkErr(err)
		if !Services.VerificationCode(res, code) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"code":    400,
				"message": "验证码错误",
			})
			return
		}
		Services.TX.Del(codeKey)
		Services.GetCode(phone)
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
			"message": CodeRequest.GetError(err.(validator.ValidationErrors)), // 注意这里要将 err 进行转换
			//"message": err.Error(), // 注意这里要将 err 进行转换
		})
		return
	}
}

//注册接口
func Register(c *gin.Context) {
	var UserRequest Requests.UserRequest
	if err := c.ShouldBind(&UserRequest); err == nil {
		name := UserRequest.Name
		password := UserRequest.Password
		age := UserRequest.Age
		email := UserRequest.Email
		phone := UserRequest.Phone
		code := UserRequest.Code
		//验证码5分钟失效
		val, err := Services.TX.Get(Services.PREFIX + phone).Result()
		if "redis: nil" != err.Error() {
			checkErr(err)
		}
		if val != code {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"code":    400,
				"message": "验证码错误或过期",
			})
			return
		}
		res := Services.DB.Create(&Models.User{Name: name, Password: password, Age: age, Email: email, Phone: phone})
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
		return
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"code":    500,
			"message": "logout error",
		})
		return
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
