package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llmclient"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)


func ProcessEmbedJob(job queue.EmbedJob){
	normalized := strings.ToLower(strings.TrimSpace(job.Prompt))

	vector, err := llmclient.GetEmbedding(normalized)
	fmt.Println("Worker:", normalized, len(vector), vector[:5])

	common.Redis.HSet(common.Ctx, "job:"+job.ID, "status", "embedding")
	common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)

	if err != nil {
		common.Redis.HSet(common.Ctx, "job:"+job.ID, map[string]interface{}{
			"status": "error",
			"error": err.Error(),
			"updated_at": time.Now().String(),
		})
		common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)
		
		log.Println("embedding failed:", err)
		return
	}

	vecJSON, _ := json.Marshal(vector)

	key := "embed:" + normalized
	err = common.Redis.Set(common.Ctx, key, vecJSON, 24*30*time.Hour).Err()
	if err != nil{
		log.Println("redis store error:", err)
        return
	}

	log.Println("Stored embedding:", key)
	queue.PushChatJob(job.ID,normalized)
}