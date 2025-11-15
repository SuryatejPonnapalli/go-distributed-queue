package users

import (
	"context"

	db "github.com/SuryatejPonnapalli/go_project/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)
type Repository struct{
	q *db.Queries
}

func NewRepository(pool *pgxpool.Pool) *Repository{

	return &Repository{q : db.New(pool)}
}

func(r *Repository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error){
	return r.q.CreateUser(ctx, arg)
}