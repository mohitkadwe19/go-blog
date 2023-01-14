package router

import "github.com/gin-gonic/gin"

// register all available route
func RegisterRouter(r *gin.Engine) {
	registerBlogRouter(r)
	registerUserRouter(r)
}
