package routes

import "github.com/gin-gonic/gin"

func SystemRoutes(router *gin.Engine) {
	system := router.Group("api/System")
	{
		system.GET("", func(c *gin.Context) {
			c.String(200, "hello-world")
		})
	}
}
