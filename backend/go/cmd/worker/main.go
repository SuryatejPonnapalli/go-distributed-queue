package main

import (
	"fmt"
	"log"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	common.InitRedis()

	fmt.Println("Worker started... listening for jobs.")

	for {
		job, err := queue.PopEmbedJob()
		if err != nil{
			log.Println("error popping job:", err)
			continue
		}

		fmt.Println("Processing job:", job.ID, "→", job.Prompt)
	}
}