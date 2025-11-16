package worker

import (
	"log"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func ProcessChatJob(job queue.EmbedJob, svc *llm.LLMService){
	log.Println("ChatJob started for:", job.Prompt)
	resp, err := svc.GetPromptResponse(job.Prompt)
	if err != nil {
		log.Println("chat failed:", err)
		return
	}

	key := "resp:" + job.Prompt

	if err := common.Redis.Set(common.Ctx, key, resp, 0).Err(); err != nil {
        log.Println("redis store error:", err)
        return
    }

	log.Println("Stored response:", key)

}