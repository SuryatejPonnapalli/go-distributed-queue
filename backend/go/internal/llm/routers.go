package llm

import (
	"github.com/gin-gonic/gin"
)

func RegisterLLMRoutes(router *gin.RouterGroup, controller *LLMController) {
	router.POST("/embed", controller.EmbedHandler)
	router.GET("/jobs/:id", controller.JobStatus)
}