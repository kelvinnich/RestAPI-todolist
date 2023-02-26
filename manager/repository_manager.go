package manager

import "authenctications/repository"

type RepositoryManagers interface {
	UserRepository() repository.UsersRepository
	TaskRepository() repository.TasksRepository
}

type repositoryManagers struct {
	infra InfraManager
}

func(r *repositoryManagers)UserRepository() repository.UsersRepository{
	return repository.NewUserRepository(r.infra.ConnecDB())
}

func(r *repositoryManagers)TaskRepository() repository.TasksRepository{
	return repository.NewTaskRepository(r.infra.ConnecDB())
}

func NewRepositoryManagers(infra InfraManager) RepositoryManagers{
	return &repositoryManagers{
		infra: infra,
	}
}