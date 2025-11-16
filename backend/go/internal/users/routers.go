package users

import "github.com/gin-gonic/gin"


func RegisterUserRoutes(router *gin.RouterGroup, controller *Controller) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}