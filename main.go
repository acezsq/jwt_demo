package main

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/jwt"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {

	r := gin.New()

	r.POST("/auth", func(c *gin.Context) {
		// 用户发送用户名和密码过来
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "无效的参数",
			})
			return
		}
		// 校验用户名和密码是否正确
		if user.Username == "q1mi" && user.Password == "q1mi123" {
			// 生成Token
			tokenString, _ := jwt.GenToken(user.Username)
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": gin.H{"token": tokenString},
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "鉴权失败",
		})
		return
	})

	r.GET("/home", jwt.JWTAuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"username": username},
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
