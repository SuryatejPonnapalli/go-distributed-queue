package main

import (
	"log"

	common "github.com/SuryatejPonnapalli/go_project/internal/common"
	"github.com/SuryatejPonnapalli/go_project/internal/users"
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

	server := gin.Default()

	//dependency injection
	userRepo := users.NewRepository(pool)
	userService := users.NewService(userRepo)
	userController := users.NewController(userService)

	//routes
	api := server.Group("/api/v1")
	users.RegisterUserRoutes(api.Group("/users"), userController)

	server.Run(":8000")
}