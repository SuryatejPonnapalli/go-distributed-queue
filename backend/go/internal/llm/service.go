package llm

import (
	"errors"

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

func (s *LLMService) checkCache(prompt string) (string, bool, error) {
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
        val, hit, _ := s.checkCache(prompt)
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
