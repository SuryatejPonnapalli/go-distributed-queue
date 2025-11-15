package common

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error){
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("CONNECTING TO:", dbURL)
	if dbURL == ""{
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err !=nil{
		return nil, err
	}

	return pool,nil
}