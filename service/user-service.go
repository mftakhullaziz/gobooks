package service

import (
	"log"

	"github.com/amifth/gorest/dto"
	"github.com/amifth/gorest/entity"
	_users "github.com/amifth/gorest/helper"
	"github.com/amifth/gorest/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID int) entity.User
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

func (service *userService) Profile(userID int) entity.User {
	return service.userRepository.FetchUserById(userID)
}

func (service *userService) AllUser() *[]_users.UsersResponse {
	users := service.userRepository.FindAll()
	usersAll := _users.NewUserArrayResponse(users)
	return &usersAll
}
