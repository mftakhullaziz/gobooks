package helper

import "github.com/amifth/apigo-gin/entity"

type UsersResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserResponse(user entity.User) UsersResponse {
	return UsersResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func NewUserArrayResponse(users []entity.User) []UsersResponse {
	uRes := []UsersResponse{}
	for _, v := range users {
		u := UsersResponse{
			ID:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		}
		uRes = append(uRes, u)
	}
	return uRes
}
