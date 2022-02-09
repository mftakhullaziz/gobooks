package service

import (
	"github.com/amifth/ApiGo/dto"
	"github.com/amifth/ApiGo/entity"
	"github.com/amifth/ApiGo/repository"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}
