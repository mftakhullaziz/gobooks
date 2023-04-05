package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/mftakhullaziz/gorest/dto"
	"github.com/mftakhullaziz/gorest/entity"
	"github.com/mftakhullaziz/gorest/helper"
	"github.com/mftakhullaziz/gorest/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID int) entity.User
	AllUser() *[]helper.UsersResponse
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

func (service *userService) AllUser() *[]helper.UsersResponse {
	users := service.userRepository.FindAll()
	usersAll := helper.NewUserArrayResponse(users)
	return &usersAll
}
