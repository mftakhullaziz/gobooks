package service

import (
	"log"

	"github.com/amifth/apigo-gin/dto"
	"github.com/amifth/apigo-gin/entity"
	_users "github.com/amifth/apigo-gin/helper"
	"github.com/amifth/apigo-gin/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	AllUser() *[]_users.UsersResponse
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) AllUser() *[]_users.UsersResponse {
	users := service.userRepository.FindAll()
	users_all := _users.NewUserArrayResponse(users)
	return &users_all
}
