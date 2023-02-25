package controller

import (
	"authenctications/usecase"
	"net/http"
	"authenctications/middleware"

	"github.com/gin-gonic/gin"
)


type UsersController interface {
	ProfileUsersController(c *gin.Context)
	FindUserByIdController(c *gin.Context)
}

type usersController struct {
	usersusecase usecase.UsersUseCase
	jwt usecase.JwtUseCase
	r *gin.Engine
}

func(u *usersController)ProfileUsersController(c *gin.Context){
	id := c.Param("id")
	user,err := u.usersusecase.ProfileUsers(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message" : "OK!", "data": user})
}

func(u *usersController)FindUserByIdController(c *gin.Context){
	id := c.Param("id")
	users, err := u.usersusecase.FindUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message" : "OK!", "data": users})
}

func NewUsersController(usersusecase usecase.UsersUseCase, jwt usecase.JwtUseCase, r *gin.Engine)UsersController{
	controler := usersController{
	usersusecase: usersusecase,
	jwt: jwt,
	r: r,
	}

	user := r.Group("/users", middleware.Authorize(jwt))
	{
		user.GET("/profile/:id", controler.ProfileUsersController)
		user.GET("/:id", controler.FindUserByIdController)
	}

	return &controler
}