package queue

import (
	"encoding/json"
	"time"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/google/uuid"
)

type EmbedJob struct {
	ID string `json:"id"`
	Prompt string `json:"prompt"`
}

func PushEmbedJob(prompt string) (string, error){
	job := EmbedJob{
		ID: uuid.New().String(),
		Prompt: prompt,
	}

	jobKey := "job:"+job.ID
	common.Redis.HSet(common.Ctx, jobKey, map[string]interface{}{
        "id": job.ID,
        "prompt": prompt,
        "status": "queued",
        "response": "",
        "error": "",
        "created_at": time.Now().String(),
        "updated_at": time.Now().String(),
    })

	data, _ := json.Marshal(job)

	err := common.Redis.LPush(common.Ctx, "embed_jobs", data).Err()
	if err != nil{
		return "", err
	}

	return job.ID, nil
}

func PopEmbedJob() (EmbedJob, error) {
	result, err := common.Redis.BRPop(common.Ctx, 0, "embed_jobs").Result()
	if err != nil {
		return EmbedJob{}, err
	}

	var job EmbedJob
	json.Unmarshal([]byte(result[1]), &job)

	return job, nil
}