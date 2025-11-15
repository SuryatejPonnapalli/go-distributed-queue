package users

import (
	"context"

	db "github.com/SuryatejPonnapalli/go_project/internal/db/generated"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{repo: repo}
}


func(s *Service) Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error){
	// hash password here
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password),10)

	//create user
	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Email: req.Email,
		Password: string(hashedPassword),
	})

	if err != nil{
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		ID: user.ID.String(),
		Email: user.Email,
	}, nil
}