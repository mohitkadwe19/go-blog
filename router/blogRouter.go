package router

import (
	"Blog/controller"

	"github.com/gin-gonic/gin"
)

func registerBlogRouter(c *gin.Engine) {

	blogRouter := c.Group("/blog")
	blogRouter.GET("/", controller.BlogGetAllController)
	blogRouter.GET("/:id", controller.BlogGetByIDController)
	blogRouter.POST("/", controller.BlogCreateController)
	blogRouter.PUT("/:id", controller.BlogUpdateController)
	blogRouter.DELETE("/:id", controller.BlogDeleteController)
}
