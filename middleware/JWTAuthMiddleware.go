package middleware

import (
	"awesomeProject/models"
	"awesomeProject/pkg/jwt_"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const ContextUserIDkey = "userID"
const ContextUser = "user"

func JwtAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		//提取token的有效部分（"Bearer "共占7位)
		tokenString = tokenString[7:]

		token, claims, err := jwt_.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		user, err := models.FindOneByUserId(userId)
		// 用户不存在
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在将user的信息写入上下文，方便读取
		ctx.Set(ContextUserIDkey, user.UserID)
		ctx.Set(ContextUser, user)

		ctx.Next()
	}
}
