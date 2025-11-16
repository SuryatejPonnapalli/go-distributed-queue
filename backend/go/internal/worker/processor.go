package worker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llmclient"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func ProcessEmbedJob(job queue.EmbedJob){
	vector, err := llmclient.GetEmbedding(job.Prompt)
	fmt.Println("Worker:", job.Prompt, len(vector), vector[:5])
	if err != nil {
		log.Println("embedding failed:", err)
		return
	}

	vecJSON, _ := json.Marshal(vector)

	key := "embed:" + job.Prompt
	err = common.Redis.Set(common.Ctx, key, vecJSON, 0).Err()
	if err != nil{
		log.Println("redis store error:", err)
        return
	}

	log.Println("Stored embedding:", key)

}