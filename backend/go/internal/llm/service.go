package llm

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
	"golang.org/x/sync/singleflight"
)

type LLMService struct {
    group singleflight.Group
}

func NewLLMService() *LLMService {
    return &LLMService{
        group: singleflight.Group{},
    }
}

func (s *LLMService) CheckCache(prompt string) (string, bool, error) {
    if prompt == "" {
        return "", false, errors.New("prompt is empty")
    }

    key := "embed:" + prompt

    val, err := common.Redis.Get(common.Ctx, key).Result()
    if err == nil {
        return val, true, nil 
    }

    return "", false, nil 
}

func (s *LLMService) FetchOrQueue(prompt string)(string, error){
    result, err, _ := s.group.Do(prompt, func() (interface{}, error) {
        val, hit, _ := s.CheckCache(prompt)
        if hit{
            return val, nil
        }

        jobID, err := queue.PushEmbedJob(prompt)
        if err != nil {
            return nil, err
        }

        return jobID, nil

    })
    return result.(string), err
}


func (s *LLMService) loadAllEmbeddings() (map[string][]float64, error){
    embeddings := make(map[string][]float64)

    cursor := uint64(0)

    for{
        keys, nextCursor, err := common.Redis.Scan(common.Ctx, cursor, "embed:*",100).Result()
        if err!= nil{
            return nil, err
        }
        cursor = nextCursor

        for _, key := range keys{
            raw, err := common.Redis.Get(common.Ctx, key).Result()
            if err != nil{
               continue
            }

            var vec []float64
            _ = json.Unmarshal([]byte(raw), &vec)

            embeddings[key] = vec
        }

        if cursor == 0{
            break
        }
    }
    return embeddings, nil
}

func (s *LLMService) FindSimilarPrompt(newVec []float64,threshold float64) (string, float64, error){
    all, err := s.loadAllEmbeddings()
    if err != nil{
        return "",0,err
    }

    bestKey := ""
    bestScore := 0.0

    for key, vec := range all{
        score := CosineSimilarity(newVec, vec)
        fmt.Println("   similarity:", score)
        if score > bestScore{
            bestScore = score
            bestKey = key
        }
    }

    if bestScore >= threshold {
        return bestKey, bestScore, nil
    }

    return "", bestScore, nil
}