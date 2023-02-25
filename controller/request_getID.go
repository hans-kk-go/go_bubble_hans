package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ID = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {

	uid, exists := c.Get(ID)
	if !exists {
		err = ErrorUserNotLogin
		return
	}
	i := uid.(int64)

	return i, err

}
