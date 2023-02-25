package controller

import (
	"awesomeProject/logic"
	"awesomeProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//1,获取参数及参数校验
	//c.shouldBandJSON
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("postHandler with invalid param", zap.Error(err))
		//请求参数有误，返回响应
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	//2,创建帖子
	err := logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}

	//3，返回响应
}

func CreatePostDetailHandler(c *gin.Context) {
	//1.获取参数(从url中获取id)
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		zap.L().Error("请求参数有误", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err,
		})
	}

	//2，根据id取出帖子数据（查数据库）
	data, err := logic.GetPostDetailHandler(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetailHandler(id) failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err,
		})
	}
	//3，返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  data,
	})
}
