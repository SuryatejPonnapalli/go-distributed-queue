package worker

import (
	"log"
	"time"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func ProcessChatJob(job queue.EmbedJob, svc *llm.LLMService){
	log.Println("ChatJob started for:", job.Prompt)

	common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
		"status":  "chatting",
		"prompt":  job.Prompt,
		"updated_at": time.Now().String(),
	})
	

	resp, err := svc.GetPromptResponse(job.Prompt)
	if err != nil {
		log.Println("chat failed:", err)

		common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
			"status": "error",
			"error": err.Error(),
			"updated_at": time.Now().String(),
		})
		
		return
	}

	key := "resp:" + job.Prompt

	if err := common.Redis.Set(common.Ctx, key, resp, 0).Err(); err != nil {
        log.Println("redis store error:", err)
        return
    }

	log.Println("Stored response:", key)

	common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
		"status": "done",
		"response": resp,
		"updated_at": time.Now().String(),
	})
	

}