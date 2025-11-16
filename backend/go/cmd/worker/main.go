package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/worker"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	common.InitRedis()

	fmt.Println("Worker started... listening for jobs.")
	fmt.Println("PYTHON_URL =", os.Getenv("PYTHON_URL"))

	for {
		job, err := queue.PopEmbedJob()
		if err != nil{
			log.Println("error popping job:", err)
			continue
		}

		worker.ProcessEmbedJob(job)
	}
}