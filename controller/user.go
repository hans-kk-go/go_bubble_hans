package controller

import (
	"awesomeProject/logic"
	"awesomeProject/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//请求参数有误，返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//手动对请求参数进行详细的业务规则校验
	if len(p.Password) == 0 || len(p.Username) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp with invalid param")
		//请求参数有误，返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//2，业务处理
	err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("service错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//3，返回响应
	c.JSON(http.StatusOK, "ok")
}

func LoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		//请求参数有误，返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//手动对请求参数进行详细的业务规则校验
	if len(p.Password) == 0 || len(p.Username) == 0 {
		zap.L().Error("SignUp with invalid param")
		//请求参数有误，返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//2.业务逻辑处理
	err2, token := logic.Login(p)
	err := err2
	if err != nil {
		zap.L().Error("service错误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败",
		})
		return
	}

	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})
}
