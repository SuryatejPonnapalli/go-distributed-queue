package main

import (
	"log"
	"os"
	"time"

	common "github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	pool, err := common.ConnectDB()
	if err != nil{
		log.Fatalf("DB connection failed: %v", err)
	}
	defer pool.Close()

	common.InitRedis()

	
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	//dependency injection

	//user
	userRepo := users.NewRepository(pool)
	userService := users.NewService(userRepo)
	userController := users.NewController(userService)

	//llm
	llmService := llm.NewLLMService()
	llmController := llm.NewLLMController(llmService)

	//routes
	api := server.Group("/api/v1")
	users.RegisterUserRoutes(api.Group("/users"), userController)
	llm.RegisterLLMRoutes(api.Group("/llm"), llmController)


	server.Run(":8000")
}