package llm

import (
	"log"
	"net/http"
	"strings"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llmclient"
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

    prompt := strings.ToLower(strings.TrimSpace(req.Prompt))

    val, hit, _ := ctl.svc.CheckCache(prompt)
    if hit {
        c.JSON(http.StatusOK, gin.H{
            "status":   "cached_exact",
            "embedding": val,
        })
        return
    }

    vec, err := llmclient.GetEmbedding(prompt)
    log.Println("Controller:", prompt, len(vec), vec[:5])
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to embed prompt"})
        return
    }

    key, score, _:= ctl.svc.FindSimilarPrompt(vec, 0.7)

    if key != ""{
        basePrompt := strings.TrimPrefix(key, "embed:")
        resp, _ := common.Redis.Get(common.Ctx, "resp:"+basePrompt).Result()

        c.JSON(http.StatusOK, gin.H{
            "status":     "semantic_reuse",
            "response":   resp,
            "similarity": score,
        })
        return
    }

    result, err := ctl.svc.FetchOrQueue(prompt)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"status": "miss", "result": result})
}