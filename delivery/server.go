package delivery

import (
	"authenctications/config"
	"authenctications/delivery/controller"
	"authenctications/manager"
	"authenctications/usecase"

	"github.com/gin-gonic/gin"
)


type Servers struct {
	usecaseManager manager.UsecaseManagers
	taskManager manager.UsecaseManagers
	engine *gin.Engine
	jwt usecase.JwtUseCase
}

func(s *Servers) Run(){
	s.initController()
	err := s.engine.Run(":5000")
	if err != nil {
		panic(err)
	}
}

func(s *Servers) initController() {
	controller.NewAuthenticationController(s.engine,s.usecaseManager.UsecaseAuth(), s.jwt)
	controller.NewUsersController(s.usecaseManager.UsersuseCase(), s.jwt, s.engine)
	controller.NewTodoListController(s.taskManager.TaskUseCase(), s.jwt, s.engine)
}

func NewServers() *Servers{
	c := config.NewConfig()
	r := gin.Default()
	infra := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManagers(infra)
	usecaseManager := manager.NewUsecaseManagers(repoManager)
	taskManager := manager.NewUsecaseManagers(repoManager)
	usecaseJWT := usecase.NewJwtUseCase()
	return &Servers{
		usecaseManager: usecaseManager,
		engine: r,
		jwt: usecaseJWT,
		taskManager: taskManager,
	}
}