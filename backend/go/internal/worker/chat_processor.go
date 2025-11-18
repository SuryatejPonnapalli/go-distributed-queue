package worker

import (
	"log"
	"strings"
	"time"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func ProcessChatJob(job queue.EmbedJob, svc *llm.LLMService){
	log.Println("ChatJob started for:", job.Prompt)

	normalized := strings.ToLower(strings.TrimSpace(job.Prompt))

	key := "resp:" + normalized
    existingResp, err := common.Redis.Get(common.Ctx, key).Result()
    if err == nil && existingResp != "" {
        log.Println("Skipping ChatJob — response already exists for:", normalized)
        return
    }


	common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
		"status":  "chatting",
		"prompt":  job.Prompt,
		"updated_at": time.Now().String(),
	})
	common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)

	resp, err := svc.GetPromptResponse(normalized)
	if err != nil {
		log.Println("chat failed:", err)

		common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
			"status": "error",
			"error": err.Error(),
			"updated_at": time.Now().String(),
		})
		common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)
		
		return
	}


	if err := common.Redis.Set(common.Ctx, key, resp, 24*7*time.Hour).Err(); err != nil {
        log.Println("redis store error:", err)
        return
    }

	log.Println("Stored response:", key)

	common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
		"status": "done",
		"response": resp,
		"updated_at": time.Now().String(),
	})
	common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)
	

}