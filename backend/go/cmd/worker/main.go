package main

import (
	"fmt"
	"os"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/worker"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	common.InitRedis()

	llmService := llm.NewLLMService()

	fmt.Println("Worker started... listening for jobs.")
	fmt.Println("PYTHON_URL =", os.Getenv("PYTHON_URL"))

	worker.StartEmbedWorkers(3)
	worker.StartChatWorkers(3, llmService)

	select {}
}