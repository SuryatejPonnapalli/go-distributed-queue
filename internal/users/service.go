package users

import (
	"context"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/auth"
	db "github.com/SuryatejPonnapalli/go-distributed-queue/internal/db/generated"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{repo: repo}
}


func(s *Service) Register(ctx context.Context, req AuthRequest) (RegisterResponse, error){
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

func(s *Service) Login(ctx context.Context, req AuthRequest) (LoginResponse, error) {
	//check if email is right and get fetched password
	user, err := s.repo.q.LoginUser(ctx, req.Email)

	if err != nil {
		return LoginResponse{}, ErrEmailNotFound
	}

	//compare passwords
	isPasswordWrong  := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if isPasswordWrong != nil{
		return LoginResponse{}, ErrWrongPassword
	}

	token, err := auth.GenerateToken(user.ID.String())
	if err != nil{
		return LoginResponse{}, ErrTokenIssue
	}

	return LoginResponse{
		Email: user.Email,
		Token: token,
	}, nil
} 