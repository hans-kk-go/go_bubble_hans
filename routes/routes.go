package routes

import (
	"awesomeProject/controller"
	"awesomeProject/logger1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger1.GinLogger(), logger1.GinRecovery(true))

	//注册业务路由
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)

	v1.POST("/login", controller.LoginHandler)
	//v1.Use(middleware.JwtAuthMiddleware())

	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "skdgj"})
	})

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.POST("/post/:id", controller.CreatePostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}
