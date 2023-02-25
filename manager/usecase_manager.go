package manager

import (
	"authenctications/usecase"
)


type UsecaseManagers interface {
	UsecaseAuth() usecase.Authentications
	UsersuseCase() usecase.UsersUseCase
}

type usecaseManagers struct {
	repoManager RepositoryManagers
}

func(u *usecaseManagers) UsecaseAuth() usecase.Authentications{
	return usecase.NewAuthentication(u.repoManager.UserRepository())
}

func(u *usecaseManagers)UsersuseCase() usecase.UsersUseCase{
	return usecase.NewUserUseCase(u.repoManager.UserRepository())
}

func NewUsecaseManagers(repositoryManagers RepositoryManagers) UsecaseManagers{
	return &usecaseManagers{repoManager: repositoryManagers}
}