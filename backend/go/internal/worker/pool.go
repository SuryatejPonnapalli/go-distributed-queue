package worker

import (
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func StartEmbedWorkers(n int){
	for i := 0;i < n; i++{
		go func(i int){
			for{
				job, err := queue.PopEmbedJob()
				if err!=nil{
					continue
				}
				ProcessEmbedJob(job)
			}
		}(i)
	}
}

func StartChatWorkers(n int, svc *llm.LLMService){
	for i := 0; i < n; i++{
		go func(i int){
			for{
				job, err:= queue.PopChatJob()
				if err!=nil{
					continue
				}
				ProcessChatJob(job, svc)
			}
		}(i)
	}
}