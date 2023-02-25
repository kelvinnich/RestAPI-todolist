package usecase

import (
	"authenctications/model"
	"authenctications/repository"
)


type UsersUseCase interface {
	ProfileUsers(id string) (model.Users, error)
	FindUserById(id string) (*model.Users,error)
}

type usersUsecase struct {
	userRepository repository.UsersRepository
}

func(u *usersUsecase)ProfileUsers(id string) (model.Users, error){
	return u.userRepository.Profile(id)
}

func(u *usersUsecase)FindUserById(id string) (*model.Users,error){
	return u.userRepository.FindUserById(id)
}

func NewUserUseCase(userrepository repository.UsersRepository) UsersUseCase{
	return &usersUsecase{
		userRepository: userrepository,
	}
}