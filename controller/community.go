package controller

import (
	"awesomeProject/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//查询到所有的社区（community_id,community_name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  data,
	})

}

func CommunityDetailHandler(c *gin.Context) {
	//获取社区id
	communityID1 := c.Param("id")
	communityID, _ := strconv.ParseInt(communityID1, 10, 64)
	//查询到所有的社区信息 以列表的形式返回
	data, err := logic.GetCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  data,
	})
}
