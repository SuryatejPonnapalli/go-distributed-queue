package llm

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LLMController struct {
    svc *LLMService
}

func NewLLMController(svc *LLMService) *LLMController {
    return &LLMController{svc}
}

func (ctl *LLMController) EmbedHandler(c *gin.Context) {
    var req EmbedRequest

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    result, err := ctl.svc.FetchOrQueue(req.Prompt)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"status": "miss", "result": result})
}