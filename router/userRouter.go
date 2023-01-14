package router

import (
	"Blog/controller"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(c *gin.Engine) {

	userRouter := c.Group("/user")
	userRouter.GET("/", controller.UserGetAllController)
	userRouter.GET("/:id", controller.UserGetByIDController)
	userRouter.POST("/", controller.UserCreateController)
	userRouter.PUT("/:id", controller.UserUpdateController)
	userRouter.DELETE("/:id", controller.UserDeleteController)
}
